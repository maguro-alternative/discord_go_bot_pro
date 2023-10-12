package botHandler

import (
	"github.com/jmoiron/sqlx"

	"github.com/maguro-alternative/discord_go_bot/db"
)

type botHandlerDB struct {
	db *db.DBHandler
}

func NewSqlDB(dbSql *sqlx.DB) *botHandlerDB {
	return &botHandlerDB{
		db: db.NewDBHandler(dbSql),
	}
}