package usecase

import (
	"sagapattern/waiter/domain"
)

type menuUsecase struct {
	repository domain.MenuRepository
}

func NewMenuUsecase(repository domain.MenuRepository) domain.MenuUsecase {
	return &menuUsecase{
		repository: repository,
	}
}

func (usecase *menuUsecase) ShowMenu() domain.Menu {
	return usecase.repository.ShowMenu()
}
