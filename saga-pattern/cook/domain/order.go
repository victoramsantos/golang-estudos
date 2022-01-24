package domain

type Order struct {
	OrderId int    `json:"order_id" validate:"required"`
	Items   []int  `json:"items" validate:"required"`
	Status  string `json:"status"`
}

type CookUsecase interface {
	Cook()
}

type BarCounterRepository interface {
	Delivery(*Order)
}

type OrderRepository interface {
	Consumes() (*Order, error)
}
