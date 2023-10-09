package botHandler

import (
	"database/sql"

	"github.com/maguro-alternative/discord_go_bot/db"
)

type botHandlerDB struct {
	db *db.DBHandler
}

func NewSqlDB(dbSql *sql.DB) *botHandlerDB {
	return &botHandlerDB{
		db: db.NewDBHandler(dbSql),
	}
}