package service

import (
	"github.com/maguro-alternative/discord_go_bot/model/envconfig"

	"github.com/bwmarrin/discordgo"
	"github.com/gorilla/sessions"
	"golang.org/x/oauth2"
)

// A DiscordService implements Discord OAuth2.
type DiscordOAuth2Service struct {
	OAuth2Config   *oauth2.Config
	CookieStore    *sessions.CookieStore
	DiscordSession *discordgo.Session
	Env            *envconfig.Env
}

// NewDiscordOAuth2Service returns new DiscordOAuth2Service.
func NewDiscordOAuth2Service(
	oauth2Config *oauth2.Config,
	cookieStore *sessions.CookieStore,
	discordSession *discordgo.Session,
	env *envconfig.Env,
) *DiscordOAuth2Service {
	return &DiscordOAuth2Service{
		OAuth2Config:   oauth2Config,
		CookieStore:    cookieStore,
		DiscordSession: discordSession,
		Env:            env,
	}
}
