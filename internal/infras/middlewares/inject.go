package middlewares

import (
	"fmt"

	"github.com/datshiro/crud/internal/infras/db"
	"github.com/datshiro/crud/internal/infras/inject"
	"github.com/labstack/echo/v4"
)

func InjectMiddleware(dbUrl string) echo.MiddlewareFunc {
	dbc, err := db.NewDB(dbUrl)
	if err != nil {
		fmt.Println(err)
		panic("Invalid database connection")
	}
	return func(next echo.HandlerFunc) echo.HandlerFunc {

		return func(e echo.Context) error {
			inject.SetDB(e, dbc)
			return next(e)
		}
	}
}
