package service

import (
	"github.com/maguro-alternative/discord_go_bot/model/envconfig"

	"github.com/jmoiron/sqlx"

	"github.com/bwmarrin/discordgo"
	"github.com/gorilla/sessions"
)

// A TODOService implements CRUD of TODO entities.
type IndexService struct {
	DB             *sqlx.DB
	CookieStore    *sessions.CookieStore
	DiscordSession *discordgo.Session
	Env            *envconfig.Env
}

// NewTODOService returns new TODOService.
func NewIndexService(
	db *sqlx.DB,
	cookieStore *sessions.CookieStore,
	discordSession *discordgo.Session,
	env *envconfig.Env,
) *IndexService {
	return &IndexService{
		DB:             db,
		CookieStore:    cookieStore,
		DiscordSession: discordSession,
		Env:            env,
	}
}
