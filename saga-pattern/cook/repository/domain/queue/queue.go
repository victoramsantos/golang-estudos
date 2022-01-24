package queue

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"sagapattern/cook/domain"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/service/sqs/sqsiface"
)

type Queue interface {
	SendMessage(order *domain.Order)
	ReadMessage() (*domain.Order, error)
}

type queue struct {
	config QueueConfig
	client sqsiface.SQSAPI
}

type QueueConfig struct {
	Region         string
	Profile        string
	QueueUrl       string
	SenderConfig   *QueueSenderConfig
	ConsumerConfig *QueueConsumerConfig
}

type QueueSenderConfig struct {
}

type QueueConsumerConfig struct {
	MaxNumberOfMessages int64
	VisibilityTimeout   int64
	WaitTimeSeconds     int64
}

func NewQueue(config *QueueConfig) *queue {
	cfg := aws.Config{
		Region:      aws.String(config.Region),
		Credentials: credentials.NewSharedCredentials("", config.Profile),
	}

	sess := session.Must(session.NewSession(&cfg))

	return &queue{
		config: *config,
		client: sqs.New(sess),
	}
}

func (queue *queue) SendMessage(order *domain.Order) {
	jsonMessage, _ := json.Marshal(order)

	sqsMessage := &sqs.SendMessageInput{
		QueueUrl:    aws.String(queue.config.QueueUrl),
		MessageBody: aws.String(string(jsonMessage)),
	}

	output, err := queue.client.SendMessage(sqsMessage)
	if err != nil {
		fmt.Printf("could not send message to queue %v: %v\n", queue.config.QueueUrl, err)
	}
	fmt.Println(output)
}

func (queue *queue) ReadMessage() (*domain.Order, error) {
	params := &sqs.ReceiveMessageInput{
		QueueUrl:            aws.String(queue.config.QueueUrl),
		MaxNumberOfMessages: aws.Int64(queue.config.ConsumerConfig.MaxNumberOfMessages),
		VisibilityTimeout:   aws.Int64(queue.config.ConsumerConfig.VisibilityTimeout),
		WaitTimeSeconds:     aws.Int64(queue.config.ConsumerConfig.WaitTimeSeconds),
	}
	resp, err := queue.client.ReceiveMessage(params)

	if err != nil {
		fmt.Println(err)
	}

	if len(resp.Messages) == 0 {
		return nil, errors.New("no message received")
	}
	message := *resp.Messages[0]
	fmt.Println(message)

	var order domain.Order
	err = json.Unmarshal([]byte(*message.Body), &order)
	if err != nil {
		log.Fatalf("Error occured during unmarshaling. Error: %s", err.Error())
	}

	return &order, nil
}
