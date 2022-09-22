package application

type productsService interface {
}

type paymentsService interface {
}

type ordersService struct {
}

func NewOrdersService() {

}

type PlaceOrderCommmand struct {
}

func (s ordersService) PlaceOrder(cmd PlaceOrderCommmand) error {

}

type MarkOrderAsPaidCommand struct {
}

func (s ordersService) MarkOrderAsPaid(cmd MarkOrderAsPaidCommand) error {

}

func (s ordersService) orderById(id orders.ID) (orsers.Order, error) {

}
