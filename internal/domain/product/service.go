package product

type Service interface {
	CreateProduct(name, description string, price float64, stock uint, categoryID uint) (*Product, error)
	GetProductByID(id uint) (*Product, error)
	GetAllProducts(page, pageSize int) ([]Product, int64, error)
	UpdateProduct(id uint, name *string, description *string, price *float64, stock *uint, categoryID *uint) (*Product, error)
	DeleteProduct(id uint) error
}
