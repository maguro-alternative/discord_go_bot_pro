package controllersDiscord

import (
	//"encoding/json"
	//"log"
	"encoding/gob"
	"net/http"

	"github.com/google/uuid"
	"golang.org/x/oauth2"

	discordModel "github.com/maguro-alternative/discord_go_bot/model/discord"
	"github.com/maguro-alternative/discord_go_bot/service"
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
	// セッションに保存する構造体の型を登録
	// これがない場合、エラーが発生する
	gob.Register(&discordModel.DiscordUser{})
	uuid := uuid.New().String()
	session, err := h.svc.CookieStore.Get(r, h.svc.Env.SessionsSecret)
	if err != nil {
		panic(err)
	}
	session.Values["state"] = uuid
	// セッションに保存
	session.Save(r, w)
	h.svc.CookieStore.Save(r, w, session)
	conf := h.svc.OAuth2Config
	// 1. 認可ページのURL
	url := conf.AuthCodeURL(uuid, oauth2.AccessTypeOffline)
	http.Redirect(w, r, url, http.StatusFound)
}
