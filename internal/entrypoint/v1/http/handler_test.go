package http_test

import (
	"example/beers/internal/beers"
	"example/beers/internal/beers/mocks"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"

	entrypoint "example/beers/internal/entrypoint/v1/http"
	"fmt"
)

var (
	BeersBrewery = []beers.Beer{
		{1, "Corona", "Modelo", "Mexico", "5", "USD"},
	}
	mockUC = []struct {
		outBeers *[]beers.Beer
		outError error
	}{
		{
			&BeersBrewery,
			nil,
		},
	}

	expect = []struct {
		in             string
		wantBody       string
		wantStatusCode int
	}{
		{
			`{}`,
			`[{"ID":1,"Name":"Corona","Brewery":"Modelo","Country":"Mexico","Price":"5","Currency":"USD"}]
`,
			http.StatusOK,
		},
	}
)

func TestSearchBeers(t *testing.T) {
	fmt.Println("BeersBrewery[1]", BeersBrewery)
	e := echo.New()

	mockUseCase := new(mocks.BeerUseCase)
	fmt.Println("mockUseCase", mockUseCase)
	for _, m := range mockUC {
		mockUseCase.On("SearchBeers").Return(m.outBeers, m.outError)
	}

	t.Run("Prueba GET con respuesta vacia ", func(t *testing.T) {

		h := entrypoint.NewServerHandler(e, mockUseCase)

		req := httptest.NewRequest(echo.GET, "/example/workshop/beers", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		h.SearchBeers(c)
		assert.Equal(t, http.StatusOK, rec.Code)
	})

	t.Run("Prueba que retorna un array de Cervezas", func(t *testing.T) {

		h := entrypoint.NewServerHandler(e, mockUseCase)
		req := httptest.NewRequest(echo.GET, "/example/workshop/beers", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		h.SearchBeers(c)

		assert.Equal(t, expect[0].wantStatusCode, rec.Code)
		assert.Equal(t, expect[0].wantBody, rec.Body.String())
	})
}
