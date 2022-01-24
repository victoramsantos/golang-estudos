package queue

//ref: https://zenidas.wordpress.com/recipes/send-message-to-sqslocalstack-with-golang/
//ref: https://onexlab-io.medium.com/localstack-sqs-a0c36fd13108

import (
	"encoding/json"
	"fmt"
	"log"
	"sagapattern/cook/domain"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/service/sqs/sqsiface"
)

type Queue interface {
	SendMessage(order *domain.Order)
	ReadMessage() domain.Order
}

type queue struct {
	url    string
	client sqsiface.SQSAPI
}

func NewQueue() *queue {
	cfg := aws.Config{
		Region:      aws.String(endpoints.UsEast1RegionID),
		Endpoint:    aws.String("http://localhost:4566"),
		Credentials: credentials.NewSharedCredentials("", "localstack"),
	}

	sess := session.Must(session.NewSession(&cfg))

	return &queue{
		url:    "http://localhost:4566/000000000000/FoodsQueue",
		client: sqs.New(sess),
	}
}

func (queue *queue) SendMessage(order *domain.Order) {
	jsonMessage, _ := json.Marshal(order)

	sqsMessage := &sqs.SendMessageInput{
		QueueUrl:    aws.String(queue.url),
		MessageBody: aws.String(string(jsonMessage)),
	}

	output, err := queue.client.SendMessage(sqsMessage)
	if err != nil {
		fmt.Printf("could not send message to queue %v: %v\n", queue.url, err)
	}
	fmt.Println(output)
}

func (queue *queue) ReadMessage() domain.Order {
	params := &sqs.ReceiveMessageInput{
		QueueUrl:            aws.String(queue.url),
		MaxNumberOfMessages: aws.Int64(1),
		VisibilityTimeout:   aws.Int64(1),
		WaitTimeSeconds:     aws.Int64(1),
	}
	resp, err := queue.client.ReceiveMessage(params)

	if err != nil {
		fmt.Println(err)
	}

	message := *resp.Messages[0]
	fmt.Println(message)

	var order domain.Order
	err = json.Unmarshal([]byte(*message.Body), &order)
	if err != nil {
		log.Fatalf("Error occured during unmarshaling. Error: %s", err.Error())
	}

	return order
}
