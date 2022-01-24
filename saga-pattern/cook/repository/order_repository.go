package repository

import (
	"sagapattern/cook/domain"
	"sagapattern/cook/repository/domain/queue"

	"github.com/spf13/viper"
)

type orderRepository struct {
	queue queue.Queue
}

func NewOrderRepository() domain.OrderRepository {
	config := queue.QueueConfig{
		Region:       viper.GetString("queues.aws_config.region"),
		Profile:      viper.GetString("queues.aws_config.profile"),
		AwsEndpoint:  viper.GetString("queues.aws_config.endpoint"),
		QueueUrl:     viper.GetString("queues.order.url"),
		SenderConfig: nil,
		ConsumerConfig: &queue.QueueConsumerConfig{
			MaxNumberOfMessages: viper.GetInt64("queues.order.max_number_of_messages"),
			VisibilityTimeout:   viper.GetInt64("queues.order.visibility_timeout"),
			WaitTimeSeconds:     viper.GetInt64("queues.order.wait_time_seconds"),
		},
	}
	return &orderRepository{
		queue: queue.NewQueue(&config),
	}
}

func (repository *orderRepository) Consumes() (*domain.Order, error) {
	order, err := repository.queue.ReadMessage()
	if err != nil {
		return nil, err
	}
	return order, nil
}
