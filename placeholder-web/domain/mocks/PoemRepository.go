package mocks

import (
	"placeholder/domain"

	"github.com/stretchr/testify/mock"
)

type PoemRepository struct {
	mock.Mock
}

func (_mock *PoemRepository) Store(poem *domain.Poem) error {
	_mock.Called(poem)

	return nil
}

func (_mock *PoemRepository) FetchPoems() []domain.Poem {
	ret := _mock.Called()

	return ret.Get(0).([]domain.Poem)
}

func (_mock *PoemRepository) FetchRandomPoem() domain.Poem {
	ret := _mock.Called()

	return ret.Get(0).(domain.Poem)
}

func (_mock *PoemRepository) FetchStaticPoem() domain.Poem {
	ret := _mock.Called()

	return ret.Get(0).(domain.Poem)
}
