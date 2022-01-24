package repository

import "sagapattern/waiter/domain"

type menuRepository struct {
	menu domain.Menu
}

func NewMenuRepository() domain.MenuRepository {
	repository := menuRepository{
		menu: domain.Menu{
			Drinks: make([]domain.Item, 0),
			Foods:  make([]domain.Item, 0),
		},
	}

	repository.menu.Foods = append(repository.menu.Foods,
		domain.Item{
			Id:   1,
			Name: "Batata Frita",
		},
	)

	repository.menu.Foods = append(repository.menu.Foods,
		domain.Item{
			Id:   2,
			Name: "Big Mac",
		},
	)

	repository.menu.Foods = append(repository.menu.Foods,
		domain.Item{
			Id:   3,
			Name: "Hot Dog",
		},
	)

	repository.menu.Foods = append(repository.menu.Foods,
		domain.Item{
			Id:   4,
			Name: "Torta de Limao",
		},
	)

	repository.menu.Foods = append(repository.menu.Foods,
		domain.Item{
			Id:   5,
			Name: "Pizza",
		},
	)

	repository.menu.Drinks = append(repository.menu.Drinks,
		domain.Item{
			Id:   1,
			Name: "Coke",
		},
	)

	repository.menu.Drinks = append(repository.menu.Drinks,
		domain.Item{
			Id:   2,
			Name: "Tequila",
		},
	)

	repository.menu.Drinks = append(repository.menu.Drinks,
		domain.Item{
			Id:   3,
			Name: "Pitu",
		},
	)

	return &repository
}

func (repository *menuRepository) ShowMenu() domain.Menu {
	return repository.menu
}
