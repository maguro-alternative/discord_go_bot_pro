package service

import (
	"golang.org/x/oauth2"
)

// A DiscordService implements Discord OAuth2.
type DiscordOAuth2Service struct {
	oauth2Config *oauth2.Config
}

// NewDiscordOAuth2Service returns new DiscordOAuth2Service.
func NewDiscordOAuth2Service(
	oauth2Config *oauth2.Config,
) *DiscordOAuth2Service {
	return &DiscordOAuth2Service{
		oauth2Config: oauth2Config,
	}
}