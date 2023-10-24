package controllersDiscord

import (
	//"log"
	"context"
	"encoding/json"
	"net/http"

	"golang.org/x/oauth2"

	"github.com/maguro-alternative/discord_go_bot/service"
	discord_model "github.com/maguro-alternative/discord_go_bot/model/discord"
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
	session, err := h.svc.SessionStore.Get(r, env.SessionsName)
	if err != nil {
		panic(err)
	}
	state, ok := session.Values["state"].(*string)
	if !ok {
		panic("state is not string")
	}
	// 2. 認可ページからリダイレクトされてきたときに送られてくるstateパラメータ
	if r.URL.Query().Get("state") != *state {
		panic("state is not match")
	}
	// 1. 認可ページのURL
	code := r.URL.Query().Get("code")
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
	ctx := context.Background()
	// 2. アクセストークンの取得
	token, err := conf.Exchange(ctx, code)
	if err != nil {
		panic(err)
	}
	session.Values["discord_access_token"] = token
	// 3. ユーザー情報の取得
	client := conf.Client(ctx, token)
	resp, err := client.Get("https://discord.com/api/users/@me")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var user discord_model.DiscordUser
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		panic(err)
	}
	// セッションに保存
	session.Values["discord_user"] = user
	//log.Println(user)
	// 4. ユーザー情報をDBに保存
	//h.svc.CreateUser(user)
	// 5. ログイン処理
	//h.svc.Login(w, r, user)
	// 6. ログイン後のページに遷移
	http.Redirect(w, r, env.FrontUrl, http.StatusFound)
}
