package service

import (
	"github.com/jmoiron/sqlx"

	"github.com/bwmarrin/discordgo"
	"github.com/gorilla/sessions"
)

// A TODOService implements CRUD of TODO entities.
type IndexService struct {
	db             *sqlx.DB
	SessionStore   *sessions.CookieStore
	DiscordSession *discordgo.Session
}

// NewTODOService returns new TODOService.
func NewIndexService(
	db *sqlx.DB,
	sessionStore *sessions.CookieStore,
	discordSession *discordgo.Session,
) *IndexService {
	return &IndexService{
		db:             db,
		SessionStore:   sessionStore,
		DiscordSession: discordSession,
	}
}
