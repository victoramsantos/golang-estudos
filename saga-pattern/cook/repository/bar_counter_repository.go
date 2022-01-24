package repository

import (
	"sagapattern/cook/domain"
	"sagapattern/cook/repository/domain/queue"

	"github.com/spf13/viper"
)

type barCounterRepository struct {
	queue queue.Queue
}

func NewBarCounterRepository() domain.BarCounterRepository {
	config := queue.QueueConfig{
		Region:         viper.GetString("queues.aws_config.region"),
		Profile:        viper.GetString("queues.aws_config.profile"),
		QueueUrl:       viper.GetString("queues.bar_counter.url"),
		SenderConfig:   &queue.QueueSenderConfig{},
		ConsumerConfig: nil,
	}
	return &barCounterRepository{
		queue: queue.NewQueue(&config),
	}
}

func (repository *barCounterRepository) Delivery(order *domain.Order) {
	repository.queue.SendMessage(order)
}
