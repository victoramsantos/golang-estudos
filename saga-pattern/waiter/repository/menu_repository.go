package repository

import (
	"sagapattern/waiter/domain"
	"sagapattern/waiter/repository/domain/database"

	"github.com/spf13/viper"
)

type menuRepository struct {
	dynamodbClient database.DynamodbClient
	cachedMenu     *domain.Menu
}

func NewMenuRepository() domain.MenuRepository {
	config := database.DynamodbConfig{
		Region:      viper.GetString("aws_config.region"),
		Profile:     viper.GetString("aws_config.profile"),
		AwsEndpoint: viper.GetString("aws_config.endpoint"),
		Table:       viper.GetString("dynamodb.table"),
	}
	return &menuRepository{
		dynamodbClient: database.NewDynamodbClient(&config),
		cachedMenu:     nil,
	}
}

func (repository *menuRepository) ShowMenu() domain.Menu {
	if repository.cachedMenu == nil {
		repository.cachedMenu = repository.dynamodbClient.GetMenu()
	}
	return *repository.cachedMenu
}
