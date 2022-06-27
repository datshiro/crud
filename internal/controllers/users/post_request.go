package users

import (
	"github.com/datshiro/crud/internal/infras/errors"
	"github.com/labstack/echo/v4"
)

func NewPostRequest() PostRequest {
	return &postRequest{}
}

type postRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (p *postRequest) GetName() string {
	return p.Name
}

func (p *postRequest) GetEmail() string {
	return p.Email
}

func (p *postRequest) Bind(e echo.Context) error {
	return e.Bind(p)
}

func (p *postRequest) Validate() error {
	if p.Name == "" {
		return errors.NewParamErr("Name must be provided")
	}
	return nil
}
