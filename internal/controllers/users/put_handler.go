package users

import (
	"net/http"

	"github.com/datshiro/crud/internal/infras/inject"
	"github.com/datshiro/crud/internal/usecases/services"
	"github.com/labstack/echo/v4"
	"github.com/volatiletech/null/v8"
)

func NewPutHandler() PutHandler {
	return &putHandler{
		NewRequest: NewPutRequest,
	}
}

type putHandler struct {
	NewRequest func() PutRequest
}

func (p *putHandler) Handle(e echo.Context) error {
	request := p.NewRequest()

	if err := request.Bind(e); err != nil {
		return err
	}

	if err := request.Validate(); err != nil {
		return err
	}

	ctx, tx := inject.CtxTx(e)
	service := services.GetUserService(ctx, tx)

	user, err := service.FindByIdForUpdate(request.GetID())
	if err != nil {
		return err
	}

	user.Name = request.GetName()
	user.Email = null.StringFrom(request.GetEmail())
	if err := service.Update(user); err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return e.JSON(http.StatusCreated, NewResponse(user))
}
