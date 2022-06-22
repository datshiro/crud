package users

import (
	"net/http"

	"github.com/datshiro/crud/internal/infras/inject"
	"github.com/datshiro/crud/internal/usecases/services"
	"github.com/labstack/echo/v4"
)

func NewPostHandler() PostHandler {
	return &postHandler{
		NewRequest: NewPostRequest,
	}
}

type postHandler struct {
	NewRequest func() PostRequest
}

func (p *postHandler) Handle(e echo.Context) error {
	request := p.NewRequest()

	if err := request.Bind(e); err != nil {
		return err
	}

	if err := request.Validate(); err != nil {
		return err
	}

	ctx, tx := inject.CtxTx(e)
	service := services.GetUserService(ctx, tx)

	user, err := service.Create(request.GetName(), request.GetEmail())
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return e.JSON(http.StatusCreated, user)
}
