package application

import "errors"

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
	price, err := price.NewPrice(cmd.ProceCents, cmd.PriceCurrency)
	if err != nil {
		return errors.Wrap(err, "Invalid product price")
	}

	p, err := products.NewProduct(products.ID(cmd.ID), cmd.Name, cmd.Description, price)

	if err != nil {
		return errors.Wrap(err, "cannoot create product")
	}

	if err := s.repo.Save(p); err != nil {
		return errors.Wrap(err, "cannot save product")
	}

	return nil
}
