package application

type productRealModel interface {
	AllProducts() ([]products.Product, err)
}

type ProductsService struct {
	repo      products.Repository
	readModel productRealModel
}

func NewProductsService() ProductsService {

}

func (s ProductsService) AllProducts() {

}

func (s ProductsService) AddProduct() error {

}
