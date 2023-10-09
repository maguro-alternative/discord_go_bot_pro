package commands

import "database/sql"

type commandHandlerDB struct {
	db *sql.DB
}

func NewSqlDB(db *sql.DB) *commandHandlerDB {
	return &commandHandlerDB{
		db: db,
	}
}