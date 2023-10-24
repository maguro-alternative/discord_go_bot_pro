package controllersDiscord

import (
	//"encoding/json"
	//"log"
	"net/http"

	"golang.org/x/oauth2"
	"github.com/google/uuid"

	"github.com/maguro-alternative/discord_go_bot/service"
	//"github.com/maguro-alternative/discord_go_bot/model"
	"github.com/maguro-alternative/discord_go_bot/model/envconfig"
)

type DiscordAuthHandler struct {
	svc *service.DiscordOAuth2Service
}

func NewDiscordAuthHandler(svc *service.DiscordOAuth2Service) *DiscordAuthHandler {
	return &DiscordAuthHandler{
		svc: svc,
	}
}

func (h *DiscordAuthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//Discordのセッションを作成
	uuid := uuid.New().String()
	env, err := envconfig.NewEnv()
	if err != nil {
		panic(err)
	}
	session, err := h.svc.SessionStore.Get(r, env.SessionsName)
	if err != nil {
		panic(err)
	}
	session.Values["state"] = uuid
	// セッションに保存
	session.Save(r, w)
	conf := &oauth2.Config{
		ClientID:     env.DiscordClientID,
		ClientSecret: env.DiscordSecret,
		Scopes:       []string{"SCOPE"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://discord.com/api/oauth2/authorize",
			TokenURL: "https://discord.com/api/oauth2/token",
		},
		RedirectURL: env.ServerUrl + "/discord/callback",
	}
	// 1. 認可ページのURL
	url := conf.AuthCodeURL(uuid, oauth2.AccessTypeOffline)
	http.Redirect(w, r, url, http.StatusFound)
}
