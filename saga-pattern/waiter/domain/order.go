package domain

type Order struct {
	OrderId int    `json:"order_id"`
	Items   []int  `json:"items" validate:"required"`
	Status  string `json:"status"`
}

type OrderUsecase interface {
	OrderRequest(*Order) error
	ListOrders() []Order
}

type OrderRepository interface {
	MakeOrder(*Order) error
	ListOrders() []Order
	LastOrderNumber() int
}
