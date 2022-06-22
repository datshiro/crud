package app

import (
	"flag"
	"os"
	"strings"

	"github.com/datshiro/crud/internal/infras/middlewares"
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
