package repository

import (
	"sagapattern/waiter/domain"
	"sagapattern/waiter/repository/domain/queue"

	"github.com/spf13/viper"
)

type orderRepository struct {
	queue queue.Queue
}

func NewOrderRepository() domain.OrderRepository {
	config := queue.QueueConfig{
		Region:         viper.GetString("queues.aws_config.region"),
		Profile:        viper.GetString("queues.aws_config.profile"),
		AwsEndpoint:    viper.GetString("queues.aws_config.endpoint"),
		QueueUrl:       viper.GetString("queues.order.url"),
		SenderConfig:   &queue.QueueSenderConfig{},
		ConsumerConfig: nil,
	}
	return &orderRepository{
		queue: queue.NewQueue(&config),
	}
}

func (repository *orderRepository) MakeOrder(order *domain.Order) error {
	repository.queue.SendMessage(order)

	return nil
}
