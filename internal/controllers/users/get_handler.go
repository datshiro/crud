package users

import (
	"net/http"

	"github.com/datshiro/crud/internal/infras/inject"
	"github.com/datshiro/crud/internal/usecases/services"
	"github.com/labstack/echo/v4"
)

func NewGetHandler() GetHandler {
	return &getHandler{
		NewRequest: NewGetRequest,
	}
}

type getHandler struct {
	NewRequest func() GetRequest
}

func (g *getHandler) Handle(c echo.Context) error {
	request := g.NewRequest()

	if err := request.Bind(c); err != nil {
		return err
	}

	if err := request.Validate(); err != nil {
		return err
	}

	ctx, exc := inject.CtxDB(c)
	service := services.GetUserService(ctx, exc)

	user, err := service.FindById(request.GetId())
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, user)
}
