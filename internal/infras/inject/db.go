package inject

import (
	"database/sql"

	"github.com/labstack/echo/v4"
)

func SetDB(e echo.Context, dbc *sql.DB) {
	e.Set(db_inject, dbc)
}

func DB(e echo.Context) *sql.DB {
	val := e.Get(db_inject)
	if val == nil {
		e.Logger().Panic("nil db")
	}

	return val.(*sql.DB)
}
