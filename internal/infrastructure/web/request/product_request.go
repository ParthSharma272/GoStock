package request

type CreateProductRequest struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" binding:"required,gt=0"`
	Stock       uint    `json:"stock" binding:"gte=0"` // gte=0 allows stock to be 0
	CategoryID  uint    `json:"category_id" binding:"required"`
}

type UpdateProductRequest struct {
	Name        *string  `json:"name,omitempty"` // Pointers to distinguish between not provided and zero value
	Description *string  `json:"description,omitempty"`
	Price       *float64 `json:"price,omitempty,gt=0"`
	Stock       *uint    `json:"stock,omitempty,gte=0"`
	CategoryID  *uint    `json:"category_id,omitempty"`
}
