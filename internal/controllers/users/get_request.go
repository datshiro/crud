package users

import (
	"github.com/datshiro/crud/internal/infras/errors"
	"github.com/labstack/echo/v4"
)

func NewGetRequest() GetRequest {
	return &getRequest{}
}

type getRequest struct {
	ID    int `param:"id"`
	Page  int `query:"page"`
	Limit int `query:"limit"`
}

func (g *getRequest) GetPage() int {
	return g.Page
}

func (g *getRequest) GetLimit() int {
	return g.Limit
}

func (g *getRequest) Bind(c echo.Context) error {
	return c.Bind(g)
}

func (g *getRequest) Validate() error {
	if g.ID == 0 {
		// if pagination not provided as well
		if g.Page == 0 || g.Limit == 0 {
			return errors.InvalidParamError
		}
		if g.Page < 0 {
			return errors.NewParamErr("page must be non negative")
		}
		if g.Limit < 0 {
			return errors.NewParamErr("limit must be non negative")
		}
	}
	if g.ID < 0 {
		return errors.InvalidIdError
	}
	return nil
}

func (g *getRequest) IsPagination() bool {
	return g.ID == 0
}

func (g *getRequest) GetId() int {
	return g.ID
}
