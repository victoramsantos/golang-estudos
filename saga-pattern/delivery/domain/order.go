package domain

type Order struct {
	OrderId string `json:"order_id" validate:"required"`
	Items   []int  `json:"items" validate:"required"`
	Status  string `json:"status"`
}

type DeliveryUsecase interface {
	Delivery()
}

type BarCounterRepository interface {
	Delivery() (*Order, error)
}
