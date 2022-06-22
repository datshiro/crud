package db

import "database/sql"

func NewDB(dbUrl string) (*sql.DB, error) {
	return sql.Open("postgres", dbUrl)
}
