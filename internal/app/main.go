package app

import (
	"github.com/datshiro/crud/internal/controllers"
	"github.com/labstack/echo/v4"
)

type App interface {
	ConfigMiddleware()
	ConfigLogLevel()
	ConfigErrorHandler()
	ConfigLogFormat()
	RegisterHandlers()
	Parse()
	Run() error
}

func NewApp() App {
	return &app{e: echo.New()}
}

type app struct {
	e         *echo.Echo
	LogLevel  string
	DbUrl     string
	Host      string
	ApiPrefix string
	Port      string
}

func (a *app) RegisterHandlers() {
	controllers.RegisterHandlers(a.e, a.ApiPrefix)
}

func (a *app) Run() error {
	address := a.Host + ":" + a.Port

	return a.e.Start(address)
}
