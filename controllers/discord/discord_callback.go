package controllersDiscord

import (
	"context"
	"encoding/json"
	"encoding/gob"
	"fmt"
	"log"
	"net/http"

	"reflect"

	discordModel "github.com/maguro-alternative/discord_go_bot/model/discord"
	"github.com/maguro-alternative/discord_go_bot/service"
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
	// セッションに保存する構造体の型を登録
	// これがない場合、エラーが発生する
	gob.Register(&discordModel.DiscordUser{})
	session, err := h.svc.CookieStore.Get(r, h.svc.Env.SessionsSecret)
	if err != nil {
		panic(err)
	}
	state, ok := session.Values["state"].(string)
	if !ok {
		fmt.Println(reflect.TypeOf(session.Values["state"]))
		panic("state is not string")
	}
	// 2. 認可ページからリダイレクトされてきたときに送られてくるstateパラメータ
	if r.URL.Query().Get("state") != state {
		session.Values["state"] = ""
		h.svc.CookieStore.Save(r, w, session)
		panic("state is not match")
	}
	session.Values["state"] = ""
	// 1. 認可ページのURL
	code := r.URL.Query().Get("code")
	conf := h.svc.OAuth2Config
	ctx := context.Background()
	// 2. アクセストークンの取得
	token, err := conf.Exchange(ctx, code)
	if err != nil {
		panic(err)
	}
	session.Values["discord_access_token"] = token.AccessToken
	// 3. ユーザー情報の取得
	client := conf.Client(ctx, token)
	resp, err := client.Get("https://discord.com/api/users/@me")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var user discordModel.DiscordUser
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		panic(err)
	}
	// セッションに保存
	session.Values["discord_user"] = user
	err = session.Save(r, w)
	if err != nil {
		panic(err)
	}
	err = h.svc.CookieStore.Save(r, w, session)
	if err != nil {
		panic(err)
	}
	log.Println(user)
	// 4. ユーザー情報をDBに保存
	//h.svc.CreateUser(user)
	// 5. ログイン処理
	//h.svc.Login(w, r, user)
	// 6. ログイン後のページに遷移
	http.Redirect(w, r, h.svc.Env.FrontUrl + "/test-user", http.StatusFound)
}
