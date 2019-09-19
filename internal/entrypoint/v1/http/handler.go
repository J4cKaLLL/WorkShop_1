package http

import (
	"net/http"

	"github.com/labstack/echo"

	beers "example/beers/internal/beers"
)

// Handler estructura que tiene las dependencias de los Handler
type Handler struct {
	usecase beers.BeerUseCase
}

// NewServerHandler cargando dependencias de caso de uso
func NewServerHandler(e *echo.Echo, usecase beers.BeerUseCase) *Handler {
	h := &Handler{
		usecase: usecase,
	}
	h.RegisterURI(e)
	return h
}

// RegisterURI Creando grupos de URI
func (h *Handler) RegisterURI(e *echo.Echo) {
	o := e.Group("/example/workshop")
	o.GET("/beers", h.SearchBeers)
}

// SearchBeers handler que expone el metodo GET /beers
func (h *Handler) SearchBeers(c echo.Context) error {
	beers, err := h.usecase.SearchBeers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}
	return c.JSON(http.StatusOK, beers)
}
