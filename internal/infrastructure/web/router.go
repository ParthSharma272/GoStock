package web

import (
	"github.com/ParthSharma272/GoStock/internal/domain/common"                 // Replace your_project_path
	"github.com/ParthSharma272/GoStock/internal/infrastructure/config"         // Replace your_project_path
	"github.com/ParthSharma272/GoStock/internal/infrastructure/web/handler"    // Replace your_project_path
	"github.com/ParthSharma272/GoStock/internal/infrastructure/web/middleware" // Replace your_project_path

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter(
	cfg *config.Config,
	authHandler *handler.AuthHandler,
	productHandler *handler.ProductHandler,
	// Add other handlers (categoryHandler, orderHandler, etc.)
) *gin.Engine {
	// gin.SetMode(gin.ReleaseMode) // Uncomment for production
	router := gin.New() // Using gin.New() for more control over middleware

	// Global Middlewares
	router.Use(gin.Logger())   // Standard Gin logger
	router.Use(gin.Recovery()) // Standard Gin recovery middleware

	// CORS Middleware
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true // For development; be more restrictive in production
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	corsConfig.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	router.Use(cors.New(corsConfig))

	// API Versioning Group
	apiV1 := router.Group("/api/v1")

	// --- Authentication Routes ---
	authRoutes := apiV1.Group("/auth")
	{
		authRoutes.POST("/register", authHandler.Register)
		authRoutes.POST("/login", authHandler.Login)
	}

	// --- Public Routes (Products, Categories) ---
	apiV1.GET("/products", productHandler.GetAllProducts)
	apiV1.GET("/products/:id", productHandler.GetProductByID)
	// apiV1.GET("/categories", categoryHandler.GetAllCategories)
	// apiV1.GET("/categories/:id", categoryHandler.GetCategoryByID)

	// --- Authenticated Routes (Common for all logged-in users) ---
	authenticated := apiV1.Group("/")
	authenticated.Use(middleware.AuthMiddleware(cfg)) // Apply JWT auth middleware
	{
		authenticated.GET("/me", authHandler.GetMyProfile) // User can see their own profile

		// Customer specific order routes (Example, implement OrderHandler)
		// customerOrderRoutes := authenticated.Group("/orders")
		// customerOrderRoutes.Use(middleware.RoleMiddleware(common.RoleCustomer))
		// {
		//  customerOrderRoutes.POST("", orderHandler.CreateOrder)
		//  customerOrderRoutes.GET("", orderHandler.GetMyOrders)
		// }
	}

	// --- Admin Routes ---
	adminRoutes := apiV1.Group("/admin")
	adminRoutes.Use(middleware.AuthMiddleware(cfg))
	adminRoutes.Use(middleware.RoleMiddleware(common.RoleAdmin))
	{
		// Product Management
		adminRoutes.POST("/products", productHandler.CreateProduct)
		adminRoutes.PUT("/products/:id", productHandler.UpdateProduct)
		adminRoutes.DELETE("/products/:id", productHandler.DeleteProduct)
		// adminRoutes.GET("/products", productHandler.GetAllProducts) // Admins can also use the public one or have a special one

		// Category Management (Example)
		// adminRoutes.POST("/categories", categoryHandler.CreateCategory)
		// adminRoutes.PUT("/categories/:id", categoryHandler.UpdateCategory)
		// adminRoutes.DELETE("/categories/:id", categoryHandler.DeleteCategory)

		// Order Management (Example)
		// adminRoutes.GET("/orders", orderHandler.GetAllOrdersAdmin)
		// adminRoutes.PUT("/orders/:id/status", orderHandler.UpdateOrderStatusAdmin)

		// User Management (Example)
		// adminRoutes.GET("/users", userHandler.GetAllUsersAdmin)
		// adminRoutes.PUT("/users/:id/role", userHandler.UpdateUserRoleAdmin)
	}

	// --- Shipper Routes (Example) ---
	// shipperRoutes := apiV1.Group("/shipper")
	// shipperRoutes.Use(middleware.AuthMiddleware(cfg))
	// shipperRoutes.Use(middleware.RoleMiddleware(common.RoleShipper, common.RoleAdmin)) // Admin can also be shipper
	// {
	//     shipperRoutes.GET("/orders/ready-to-ship", orderHandler.GetOrdersForShipper)
	//     shipperRoutes.PUT("/orders/:id/ship", orderHandler.ShipOrder)
	// }

	return router
}
