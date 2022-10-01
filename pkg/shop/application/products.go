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

type AddProductCommand struct {
	ID            string
	Name          string
	Description   string
	ProceCents    uint
	PriceCurrency uint
}

func (s ProductsService) AddProduct(cmd AddProductCommand) error {
	price.NewPrice(cmd.ProceCents, cmd.PriceCurrency)

	products.NewProduct(products.ID(cmd.ID), cmd.Name, cmd.Description, price)

	s.repo.Save
}
