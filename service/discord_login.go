package service

import (
	"github.com/gorilla/sessions"
	"golang.org/x/oauth2"
)

// A DiscordService implements Discord OAuth2.
type DiscordOAuth2Service struct {
	oauth2Config *oauth2.Config
	SessionStore *sessions.CookieStore
}

// NewDiscordOAuth2Service returns new DiscordOAuth2Service.
func NewDiscordOAuth2Service(
	oauth2Config *oauth2.Config,
	sessionStore *sessions.CookieStore,
) *DiscordOAuth2Service {
	return &DiscordOAuth2Service{
		oauth2Config: oauth2Config,
		SessionStore: sessionStore,
	}
}
