package commands

import (
	"database/sql"

	"github.com/maguro-alternative/discord_go_bot/db"
)


type commandHandlerDB struct {
	db *db.DBHandler
}

func NewSqlDB(dbSql *sql.DB) *commandHandlerDB {
	return &commandHandlerDB{
		db: db.NewDBHandler(dbSql),
	}
}