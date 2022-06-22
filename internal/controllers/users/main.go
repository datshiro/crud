package users

import "github.com/labstack/echo/v4"

func RegisterHandlers(e *echo.Echo, apiPrefix string) {
	e.POST(apiPrefix, NewPostHandler().Handle)
}
