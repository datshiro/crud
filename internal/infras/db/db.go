package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func NewDB(dbUrl string) (*sql.DB, error) {
	return sql.Open("postgres", dbUrl)
}
