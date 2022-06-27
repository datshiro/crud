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
	if request.IsPagination() {
		return g.getPagination(request, c)
	}

	ctx, exc := inject.CtxDB(c)
	service := services.GetUserService(ctx, exc)

	user, err := service.FindById(request.GetId())
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, NewResponse(user))
}

func (g *getHandler) getPagination(request GetRequest, c echo.Context) error {
	ctx, exec := inject.CtxDB(c)
	service := services.GetUserService(ctx, exec)
	users, err := service.GetPagination(request.GetPage(), request.GetLimit())
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, NewResponses(users))
}
