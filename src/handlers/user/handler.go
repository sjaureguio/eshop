package user

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/sjaureguio/eshop/internal/user/application"
	"github.com/sjaureguio/eshop/internal/user/domain"
)

type handler struct {
	useCase application.WebHandler
}

func (h handler) new(uc application.UseCase) *handler {
	return &handler{useCase: uc}
}

func (h handler) Create(c echo.Context) error {
	m := domain.User{}

	if err := c.Bind(&m); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := h.useCase.Create(&m); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, m)
}

func (h handler) GetAll(c echo.Context) error {
	users, err := h.useCase.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, users)
}
