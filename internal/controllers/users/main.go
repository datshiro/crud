package users

import (
	"path"

	"github.com/labstack/echo/v4"
)

func RegisterHandlers(e *echo.Echo, apiPrefix string) {
	e.POST(apiPrefix, NewPostHandler().Handle)
	e.PUT(apiPrefix, NewPutHandler().Handle)
	e.GET(apiPrefix, NewGetHandler().Handle)

	e.GET(path.Join(apiPrefix, ":id"), NewGetHandler().Handle)
	e.DELETE(path.Join(apiPrefix, ":id"), NewDeleteHandler().Handle)
}
