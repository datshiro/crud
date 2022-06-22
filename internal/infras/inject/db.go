package inject

import (
	"context"
	"database/sql"

	"github.com/labstack/echo/v4"
)

func Ctx(_ echo.Context) context.Context {
	return context.Background()
}

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

func CtxDB(e echo.Context) (context.Context, *sql.DB) {
	ctx := Ctx(e)
	return ctx, DB(e)
}

func CtxTx(e echo.Context) (context.Context, *sql.Tx) {
	ctx, db := CtxDB(e)
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return nil, nil
	}
	return ctx, tx
}
