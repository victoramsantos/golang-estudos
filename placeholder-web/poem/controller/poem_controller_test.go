package controller_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"placeholder/domain"
	"placeholder/domain/mocks"
	"placeholder/poem/controller"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestStoreUsecase(t *testing.T) {
	mockPoemUsecase := new(mocks.PoemUsecase)

	t.Run("success", func(t *testing.T) {
		mockPoemUsecase.On("Store", mock.AnythingOfType("*domain.Poem")).Return(nil).Once()

		poem := domain.Poem{
			Content: "some-content",
			Author:  "some-author",
		}

		j, err := json.Marshal(poem)
		assert.NoError(t, err)

		e := echo.New()
		req, err := http.NewRequest(echo.POST, "/poem", strings.NewReader(string(j)))
		assert.NoError(t, err)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/poem")

		controller := controller.PoemController{
			PoemUsecase: mockPoemUsecase,
		}
		err = controller.Store(c)
		require.NoError(t, err)

		assert.Equal(t, http.StatusCreated, rec.Code)
		mockPoemUsecase.AssertExpectations(t)
	})
}
