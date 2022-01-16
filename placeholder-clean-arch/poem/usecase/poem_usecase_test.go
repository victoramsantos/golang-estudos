package usecase_test

import (
	"placeholder/domain"
	"placeholder/domain/mocks"
	"placeholder/poem/usecase"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFetchStaticUsecase(t *testing.T) {
	mockPoemRepo := new(mocks.PoemRepository)
	mockPoem := domain.Poem{
		Content: "some-content",
		Author:  "some-author",
	}

	t.Run("success", func(t *testing.T) {
		mockPoemRepo.On("FetchStaticPoem").Return(mockPoem).Once()

		u := usecase.NewPoemUsecase(mockPoemRepo)
		poem := u.FetchStaticUsecase()
		assert.Equal(t, mockPoem, poem)

		mockPoemRepo.AssertExpectations(t)
	})
}

func TestFetchRandomUsecase(t *testing.T) {
	mockPoemRepo := new(mocks.PoemRepository)
	mockPoem := domain.Poem{
		Content: "some-content",
		Author:  "some-author",
	}

	t.Run("success", func(t *testing.T) {
		mockPoemRepo.On("FetchRandomPoem").Return(mockPoem).Once()

		u := usecase.NewPoemUsecase(mockPoemRepo)
		poem := u.FetchRandomUsecase()
		assert.Equal(t, mockPoem, poem)

		mockPoemRepo.AssertExpectations(t)
	})
}
