package service

import (
	"github.com/jmoiron/sqlx"

	"github.com/bwmarrin/discordgo"
)

// A TODOService implements CRUD of TODO entities.
type IndexService struct {
	db             *sqlx.DB
	DiscordSession *discordgo.Session
}

// NewTODOService returns new TODOService.
func NewIndexService(db *sqlx.DB, discordSession *discordgo.Session) *IndexService {
	return &IndexService{
		db:             db,
		DiscordSession: discordSession,
	}
}
