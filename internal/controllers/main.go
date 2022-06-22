package controllers

import (
	"path"

	"github.com/datshiro/crud/internal/controllers/users"
	"github.com/labstack/echo/v4"
)

const (
	UserPath = "users"
)

func RegisterHandlers(e *echo.Echo, apiPrefix string) {
	users.RegisterHandlers(e, path.Join(apiPrefix, UserPath))
}
