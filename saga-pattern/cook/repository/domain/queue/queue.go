package queue

import (
	"encoding/json"
	"errors"
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
	AwsEndpoint    string
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
		Endpoint:    aws.String(config.AwsEndpoint),
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

	_, err := queue.client.SendMessage(sqsMessage)
	if err != nil {
		log.Printf("could not send message to queue %v: %v\n", queue.config.QueueUrl, err)
	}
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
		log.Println(err)
	}

	if len(resp.Messages) == 0 {
		log.Println("No message received")
		return nil, errors.New("no message received")
	}
	message := *resp.Messages[0]

	var order domain.Order
	err = json.Unmarshal([]byte(*message.Body), &order)
	if err != nil {
		log.Printf("Error occured during unmarshaling. Error: %s\n", err.Error())
		return nil, errors.New("unmarshaling error")
	}

	receiptHandle := resp.Messages[0].ReceiptHandle

	_, err = queue.client.DeleteMessage(&sqs.DeleteMessageInput{
		QueueUrl:      aws.String(queue.config.QueueUrl),
		ReceiptHandle: receiptHandle,
	})

	if err != nil {
		log.Printf("Got an error while trying to delete message: %v", err)
		return nil, errors.New("got error while trying to delete message")
	}

	return &order, nil
}
