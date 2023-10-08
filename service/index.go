package service

import (
	"database/sql"

	"github.com/bwmarrin/discordgo"
)

// A TODOService implements CRUD of TODO entities.
type IndexService struct {
	db *sql.DB
	DiscordSession *discordgo.Session
}

// NewTODOService returns new TODOService.
func NewIndexService(db *sql.DB,discordSession *discordgo.Session) *IndexService {
	return &IndexService{
		db: db,
		DiscordSession : discordSession,
	}
}
