package usecase_test

import (
	"placeholder/domain"
	"placeholder/domain/mocks"
	"placeholder/poem/usecase"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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

func TestFetchUsecase(t *testing.T) {
	mockPoemRepo := new(mocks.PoemRepository)
	mockPoems := make([]domain.Poem, 0)
	mockPoems = append(mockPoems, domain.Poem{
		Content: "some-content",
		Author:  "some-author",
	})

	t.Run("success", func(t *testing.T) {
		mockPoemRepo.On("FetchPoems").Return(mockPoems).Once()

		u := usecase.NewPoemUsecase(mockPoemRepo)
		poems := u.FetchUsecase()
		assert.NotEmpty(t, poems)
		assert.Len(t, poems, len(mockPoems))
		assert.Equal(t, mockPoems[0], poems[0])
		mockPoemRepo.AssertExpectations(t)
	})
}

func TestStoreUsecase(t *testing.T) {
	mockPoemRepo := new(mocks.PoemRepository)

	t.Run("success", func(t *testing.T) {
		mockPoemRepo.On("Store", mock.AnythingOfType("*domain.Poem")).Return(nil).Once()

		poem := domain.Poem{
			Content: "some-content",
			Author:  "some-author",
		}

		u := usecase.NewPoemUsecase(mockPoemRepo)
		err := u.Store(&poem)
		assert.NoError(t, err)
		mockPoemRepo.AssertExpectations(t)
	})
}
