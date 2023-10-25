package testRouter

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	discordModel "github.com/maguro-alternative/discord_go_bot/model/discord"
	"github.com/maguro-alternative/discord_go_bot/service"
)

type IndexHandler struct {
	svc *service.IndexService
}

// NewTODOHandler returns TODOHandler based http.Handler.
func NewAuthCheckHandler(svc *service.IndexService) *IndexHandler {
	return &IndexHandler{
		svc: svc,
	}
}

func (h *IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	session, err := h.svc.CookieStore.Get(r, h.svc.Env.SessionsSecret)
	if err != nil {
		panic(err)
	}
	fmt.Println(session.Values["discord_user"])
	// セッションに保存されているdiscorduserを取得
	discordUser, ok := session.Values["discord_user"].(*discordModel.DiscordUser)
	if !ok {
		panic("discorduser is not string")
	}
	err = json.NewEncoder(w).Encode(&discordUser)
	if err != nil {
		log.Println(err)
	}
}
