package users

import (
	"github.com/datshiro/crud/internal/infras/errors"
	"github.com/labstack/echo/v4"
)

func NewDeleteRequest() DeleteRequest {
	return &deleteRequest{}
}

type deleteRequest struct {
	ID int `param:"id"`
}

func (g *deleteRequest) Bind(c echo.Context) error {
	return c.Bind(g)
}

func (g *deleteRequest) Validate() error {
	if g.ID <= 0 {
		return errors.InvalidIdError
	}
	return nil
}

func (g *deleteRequest) GetId() int {
	return g.ID
}
