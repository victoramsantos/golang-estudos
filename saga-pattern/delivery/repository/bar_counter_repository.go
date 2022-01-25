package repository

import (
	"sagapattern/delivery/domain"
	"sagapattern/delivery/repository/domain/queue"

	"github.com/spf13/viper"
)

type barCounterRepository struct {
	queue queue.Queue
}

func NewBarCounterRepository() domain.BarCounterRepository {
	config := queue.QueueConfig{
		Region:       viper.GetString("aws_config.region"),
		Profile:      viper.GetString("aws_config.profile"),
		AwsEndpoint:  viper.GetString("aws_config.endpoint"),
		QueueUrl:     viper.GetString("queues.bar_counter.url"),
		SenderConfig: nil,
		ConsumerConfig: &queue.QueueConsumerConfig{
			MaxNumberOfMessages: viper.GetInt64("queues.bar_counter.max_number_of_messages"),
			VisibilityTimeout:   viper.GetInt64("queues.bar_counter.visibility_timeout"),
			WaitTimeSeconds:     viper.GetInt64("queues.bar_counter.wait_time_seconds"),
		},
	}
	return &barCounterRepository{
		queue: queue.NewQueue(&config),
	}
}

func (repository *barCounterRepository) Delivery() (*domain.Order, error) {
	order, err := repository.queue.ReadMessage()
	if err != nil {
		return nil, err
	}
	return order, nil
}
