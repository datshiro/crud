package app

import (
	"flag"
	"net/http"
	"os"
	"strings"

	"github.com/datshiro/crud/internal/infras/errors"
	"github.com/datshiro/crud/internal/infras/middlewares"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

var (
	DefaultDBUrl     = "postgres://postgres:postgres@localhost:5432/crud?sslmode=disable"
	DefaultPort      = "8080"
	DefaultHost      = "localhost"
	DefaultLogLevel  = "error"
	DefaultApiPrefix = "/api"

	DBUrlEnv     = "DB_URL"
	PortEnv      = "SERVER_PORT"
	HostEnv      = "SERVER_HOST"
	LogLevelEnv  = "LOG_LEVEL"
	ApiPrefixEnv = "API_PREFIX"
)

func GetEnv(key string, defaultValue string) string {
	val := os.Getenv(key)
	if val == "" {
		return defaultValue
	}
	return val
}

func (a *app) Parse() {
	flag.StringVar(&a.DbUrl, "db", GetEnv(DBUrlEnv, DefaultDBUrl), "DB Url for DB connection")
	flag.StringVar(&a.LogLevel, "log", GetEnv(LogLevelEnv, DefaultLogLevel), "Log level")
	flag.StringVar(&a.Port, "port", GetEnv(PortEnv, DefaultPort), "Running port")
	flag.StringVar(&a.Host, "host", GetEnv(HostEnv, DefaultHost), "Running host")
	flag.StringVar(&a.ApiPrefix, "api", GetEnv(ApiPrefixEnv, DefaultApiPrefix), "Server API Prefix")
	flag.Parse()
}

func (a *app) ConfigMiddleware() {
	a.e.Use(middleware.Recover())

	a.e.Use(middlewares.InjectMiddleware(a.DbUrl))
}

func (a *app) ConfigLogLevel() {
	logger := strings.ToLower(a.LogLevel)
	switch logger {
	case "error":
		a.e.Logger.SetLevel(log.ERROR)
	case "info":
		a.e.Logger.SetLevel(log.INFO)
	case "warn":
		a.e.Logger.SetLevel(log.WARN)
	case "debug":
		fallthrough
	default:
		a.e.Logger.SetLevel(log.DEBUG)
	}
}
func (a *app) ConfigErrorHandler() {
	defaultHandler := a.e.HTTPErrorHandler
	a.e.HTTPErrorHandler = func(err error, c echo.Context) {
		if c.IsWebSocket() {
			return
		}
		switch err.(type) {
		case errors.CustomError, errors.CustomParamError:
			_ = c.JSON(http.StatusBadRequest, NewErrorResponse(err))

		default:
			defaultHandler(err, c)
		}
	}
}

// configLogHeader change echo global log format, and adhoc log prefix
// for more readable. The default one produces log in JSON format, with is
// intended to be collected by other tools, but we're not using such tools yet.
func (a *app) ConfigLogFormat() {
	a.e.HideBanner = true
	a.e.Logger.SetOutput(os.Stderr)

	// make echo context log more readable.
	if l, ok := a.e.Logger.(*log.Logger); ok {
		l.SetHeader("${time_rfc3339} ${level} ${short_file}:${line}")
	}

	// make echo request/response log (once per request) more readable
	a.e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `${time_rfc3339} requestID=${id} remote_ip=${remote_ip} ` +
			`${method} ${path} status=${status} err=${error} ` +
			`latency=${latency_human} user_agent=${user_agent}`,
	}))
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func NewErrorResponse(err error) ErrorResponse {
	return ErrorResponse{
		Error: err.Error(),
	}
}
