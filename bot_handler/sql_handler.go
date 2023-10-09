package botHandler

import (
	"database/sql"
)

type botHandlerDB struct {
	db *sql.DB
}

func NewSqlDB(db *sql.DB) *botHandlerDB {
	return &botHandlerDB{
		db: db,
	}
}