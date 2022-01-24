package database

import (
	"sagapattern/waiter/domain"

	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

type DynamodbClient interface {
	GetMenu() *domain.Menu
}

type dynamodbClient struct {
	config DynamodbConfig
	client *dynamodb.DynamoDB
}

type DynamodbConfig struct {
	Region      string
	Profile     string
	AwsEndpoint string
	Table       string
}

func NewDynamodbClient(config *DynamodbConfig) *dynamodbClient {
	cfg := aws.Config{
		Region:      aws.String(config.Region),
		Endpoint:    aws.String(config.AwsEndpoint),
		Credentials: credentials.NewSharedCredentials("", config.Profile),
	}

	sess := session.Must(session.NewSession(&cfg))

	return &dynamodbClient{
		config: *config,
		client: dynamodb.New(sess),
	}
}

func (dynamodbClient *dynamodbClient) GetMenu() *domain.Menu {
	return &domain.Menu{
		Foods:  dynamodbClient.getItemsByCategory("FOOD"),
		Drinks: dynamodbClient.getItemsByCategory("DRINK"),
	}
}

func (dynamodbClient *dynamodbClient) getItemsByCategory(category string) []domain.Item {
	filt := expression.Name("Category").Equal(expression.Value(category))
	expr, err := expression.NewBuilder().WithFilter(filt).Build()
	if err != nil {
		log.Printf("failed to create the Expression, %v\n", err)
		return nil
	}
	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 aws.String(dynamodbClient.config.Table),
		Limit:                     aws.Int64(20),
	}

	result, err := dynamodbClient.client.Scan(params)
	if err != nil {
		log.Printf("failed to make Query API call, %v\n", err)
		return nil
	}

	items := make([]domain.Item, 0)

	for _, i := range result.Items {
		item := domain.Item{}
		err = dynamodbattribute.UnmarshalMap(i, &item)

		if err != nil {
			log.Fatalf("Got error unmarshalling: %s", err)
		}
		items = append(items, item)
	}
	return items
}
