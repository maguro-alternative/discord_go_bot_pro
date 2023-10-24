package service

import (
	"github.com/jmoiron/sqlx"

	"github.com/bwmarrin/discordgo"
	"github.com/gorilla/sessions"
)

// A TODOService implements CRUD of TODO entities.
type IndexService struct {
	db             *sqlx.DB
	Session        *sessions.Session
	DiscordSession *discordgo.Session
}

// NewTODOService returns new TODOService.
func NewIndexService(
	db *sqlx.DB,
	session *sessions.Session,
	discordSession *discordgo.Session,
) *IndexService {
	return &IndexService{
		db:             db,
		Session:        session,
		DiscordSession: discordSession,
	}
}
