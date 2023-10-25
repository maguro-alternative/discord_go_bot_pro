package router

import (
	"net/http"

	"github.com/jmoiron/sqlx"
	"golang.org/x/oauth2"

	"github.com/maguro-alternative/discord_go_bot/model/envconfig"
	serverHandler "github.com/maguro-alternative/discord_go_bot/server_handler"
	"github.com/maguro-alternative/discord_go_bot/server_handler/middleware"
	"github.com/maguro-alternative/discord_go_bot/service"
	controllersDiscord "github.com/maguro-alternative/discord_go_bot/controllers/discord"

	testRouter "github.com/maguro-alternative/discord_go_bot/server_handler/_test"

	"github.com/bwmarrin/discordgo"
	"github.com/gorilla/sessions"
	"github.com/justinas/alice"
)

func NewRouter(
	indexDB *sqlx.DB,
	cookieStore *sessions.CookieStore,
	discordSession *discordgo.Session,
	env *envconfig.Env,
) *http.ServeMux {
	conf := &oauth2.Config{
		ClientID:     env.DiscordClientID,
		ClientSecret: env.DiscordSecret,
		Scopes:       []string{"identify"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://discord.com/api/oauth2/authorize",
			TokenURL: "https://discord.com/api/oauth2/token",
		},
		RedirectURL: env.ServerUrl + "/discord-callback/",
	}
	// create a *service.TODOService type variable using the *sql.DB type variable
	var indexService = service.NewIndexService(
		indexDB,
		cookieStore,
		discordSession,
		env,
	)
	var discordOAuth2Service = service.NewDiscordOAuth2Service(
		conf,
		cookieStore,
		discordSession,
		env,
	)

	// register routes
	mux := http.NewServeMux()
	middleChain := alice.New(middleware.CORS)
	mux.Handle("/", middleChain.Then(serverHandler.NewIndexHandler(indexService)))
	mux.Handle("/discord-auth-check", middleChain.Then(testRouter.NewAuthCheckHandler(indexService)))
	mux.Handle("/discord/auth", middleChain.Then(controllersDiscord.NewDiscordAuthHandler(discordOAuth2Service)))
	mux.Handle("/discord-callback/", middleChain.Then(controllersDiscord.NewDiscordCallbackHandler(discordOAuth2Service)))
	return mux
}
