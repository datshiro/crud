package users

import (
	"github.com/datshiro/crud/internal/infras/errors"
	"github.com/labstack/echo/v4"
)

func NewPutRequest() PutRequest {
	return &putRequest{}
}

type putRequest struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (p *putRequest) GetName() string {
	return p.Name
}

func (p *putRequest) GetEmail() string {
	return p.Email
}

func (p *putRequest) GetID() int {
	return p.ID
}

func (p *putRequest) Bind(e echo.Context) error {
	return e.Bind(p)
}

func (p *putRequest) Validate() error {
	if p.ID <= 0 {
		return errors.InvalidIdError
	}
	if p.Name == "" {
		return errors.NewParamErr("Name must be provided")
	}
	return nil
}
