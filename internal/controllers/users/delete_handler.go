package users

import (
	"net/http"

	"github.com/datshiro/crud/internal/infras/inject"
	"github.com/datshiro/crud/internal/usecases/services"
	"github.com/labstack/echo/v4"
)

func NewDeleteHandler() DeleteHandler {
	return &deleteHandler{
		NewRequest: NewDeleteRequest,
	}
}

type deleteHandler struct {
	NewRequest func() DeleteRequest
}

func (g *deleteHandler) Handle(c echo.Context) error {
	request := g.NewRequest()

	if err := request.Bind(c); err != nil {
		return err
	}

	if err := request.Validate(); err != nil {
		return err
	}

	ctx, exc := inject.CtxDB(c)
	service := services.GetUserService(ctx, exc)

	err := service.DeleteById(request.GetId())
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, nil)
}
