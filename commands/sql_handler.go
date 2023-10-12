package commands

import (
	"github.com/jmoiron/sqlx"

	"github.com/maguro-alternative/discord_go_bot/db"
)


type commandHandlerDB struct {
	db *db.DBHandler
}

func NewSqlDB(dbSql *sqlx.DB) *commandHandlerDB {
	return &commandHandlerDB{
		db: db.NewDBHandler(dbSql),
	}
}