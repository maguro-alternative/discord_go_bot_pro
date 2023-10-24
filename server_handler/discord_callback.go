package serverHandler

import (
	//"encoding/json"
	//"log"
	"net/http"

	"golang.org/x/oauth2"

	"github.com/maguro-alternative/discord_go_bot/service"
	//"github.com/maguro-alternative/discord_go_bot/model"
	"github.com/maguro-alternative/discord_go_bot/model/envconfig"
)

type DiscordCallbackHandler struct {
	svc *service.DiscordOAuth2Service
}

func NewDiscordCallbackHandler(svc *service.DiscordOAuth2Service) *DiscordCallbackHandler {
	return &DiscordCallbackHandler{
		svc: svc,
	}
}

func (h *DiscordCallbackHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//Discordのセッションを作成
	env, err := envconfig.NewEnv()
	if err != nil {
		panic(err)
	}
	conf := &oauth2.Config{
		ClientID:     env.DiscordClientID,
		ClientSecret: env.DiscordSecret,
		Scopes:       []string{"SCOPE"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https//discord.com/api/oauth2/authorize",
			TokenURL: "https://discord.com/api/oauth2/token",
		},
	}
	// 1. 認可ページのURL
	url := conf.AuthCodeURL("", oauth2.AccessTypeOffline)
	http.Redirect(w, r, url, http.StatusFound)
}
