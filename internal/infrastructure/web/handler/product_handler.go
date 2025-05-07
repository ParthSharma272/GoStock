package handler

import (
	domainProduct "github.com/ParthSharma272/GoStock/internal/domain/product" // Replace your_project_path
	"github.com/ParthSharma272/GoStock/internal/infrastructure/web/request"   // Replace your_project_path
	"github.com/ParthSharma272/GoStock/internal/infrastructure/web/response"  // Replace your_project_path
	"net/http"
	"strconv"

	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProductHandler struct {
	productService domainProduct.Service
}

func NewProductHandler(productService domainProduct.Service) *ProductHandler {
	return &ProductHandler{productService: productService}
}

// CreateProduct godoc
// @Summary Create a new product
// @Description Create a new product (Admin only)
// @Tags products
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer Token"
// @Param product body request.CreateProductRequest true "Product data"
// @Success 201 {object} domainProduct.Product
// @Failure 400 {object} response.StandardErrorResponse
// @Failure 401 {object} response.StandardErrorResponse
// @Failure 403 {object} response.StandardErrorResponse
// @Failure 500 {object} response.StandardErrorResponse
// @Router /admin/products [post]
func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var req request.CreateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.StandardErrorResponse{Error: err.Error()})
		return
	}

	product, err := h.productService.CreateProduct(req.Name, req.Description, req.Price, req.Stock, req.CategoryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.StandardErrorResponse{Error: "Failed to create product: " + err.Error()})
		return
	}
	c.JSON(http.StatusCreated, product)
}

// GetProductByID godoc
// @Summary Get a product by ID
// @Description Get details of a specific product
// @Tags products
// @Produce  json
// @Param id path int true "Product ID"
// @Success 200 {object} domainProduct.Product
// @Failure 400 {object} response.StandardErrorResponse
// @Failure 404 {object} response.StandardErrorResponse
// @Router /products/{id} [get]
func (h *ProductHandler) GetProductByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.StandardErrorResponse{Error: "Invalid product ID format"})
		return
	}

	product, err := h.productService.GetProductByID(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, response.StandardErrorResponse{Error: "Product not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, response.StandardErrorResponse{Error: "Failed to retrieve product: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, product)
}

// GetAllProducts godoc
// @Summary Get all products (paginated)
// @Description Get a list of all available products with pagination
// @Tags products
// @Produce  json
// @Param page query int false "Page number (default: 1)"
// @Param pageSize query int false "Number of items per page (default: 10)"
// @Success 200 {object} response.PaginatedResponse{Data=[]domainProduct.Product}
// @Failure 500 {object} response.StandardErrorResponse
// @Router /products [get]
func (h *ProductHandler) GetAllProducts(c *gin.Context) {
	pageQuery := c.DefaultQuery("page", "1")
	pageSizeQuery := c.DefaultQuery("pageSize", "10")

	page, _ := strconv.Atoi(pageQuery)
	pageSize, _ := strconv.Atoi(pageSizeQuery)

	products, total, err := h.productService.GetAllProducts(page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.StandardErrorResponse{Error: "Failed to retrieve products: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.PaginatedResponse{
		Data:     products,
		Total:    total,
		Page:     page,
		PageSize: pageSize,
		LastPage: (total + int64(pageSize) - 1) / int64(pageSize),
	})
}

// UpdateProduct godoc
// @Summary Update an existing product
// @Description Update details of an existing product (Admin only)
// @Tags products
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer Token"
// @Param id path int true "Product ID"
// @Param product body request.UpdateProductRequest true "Product data to update"
// @Success 200 {object} domainProduct.Product
// @Failure 400 {object} response.StandardErrorResponse
// @Failure 401 {object} response.StandardErrorResponse
// @Failure 403 {object} response.StandardErrorResponse
// @Failure 404 {object} response.StandardErrorResponse
// @Failure 500 {object} response.StandardErrorResponse
// @Router /admin/products/{id} [put]
func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.StandardErrorResponse{Error: "Invalid product ID format"})
		return
	}

	var req request.UpdateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.StandardErrorResponse{Error: err.Error()})
		return
	}

	product, err := h.productService.UpdateProduct(uint(id), req.Name, req.Description, req.Price, req.Stock, req.CategoryID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, response.StandardErrorResponse{Error: "Product not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, response.StandardErrorResponse{Error: "Failed to update product: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, product)
}

// DeleteProduct godoc
// @Summary Delete a product
// @Description Delete a product by ID (Admin only)
// @Tags products
// @Produce  json
// @Param Authorization header string true "Bearer Token"
// @Param id path int true "Product ID"
// @Success 200 {object} response.StandardSuccessResponse
// @Failure 400 {object} response.StandardErrorResponse
// @Failure 401 {object} response.StandardErrorResponse
// @Failure 403 {object} response.StandardErrorResponse
// @Failure 404 {object} response.StandardErrorResponse
// @Failure 500 {object} response.StandardErrorResponse
// @Router /admin/products/{id} [delete]
func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.StandardErrorResponse{Error: "Invalid product ID format"})
		return
	}

	err = h.productService.DeleteProduct(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) { // Or a custom error from service
			c.JSON(http.StatusNotFound, response.StandardErrorResponse{Error: "Product not found or cannot be deleted"})
			return
		}
		c.JSON(http.StatusInternalServerError, response.StandardErrorResponse{Error: "Failed to delete product: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, response.StandardSuccessResponse{Message: "Product deleted successfully"})
}
