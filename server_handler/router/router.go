package router

import (
	"net/http"

	"github.com/jmoiron/sqlx"

	serverHandler "github.com/maguro-alternative/discord_go_bot/server_handler"
	"github.com/maguro-alternative/discord_go_bot/service"
	"github.com/maguro-alternative/discord_go_bot/server_handler/middleware"

	"github.com/bwmarrin/discordgo"
	"github.com/gorilla/sessions"
	"github.com/justinas/alice"
)

func NewRouter(
	indexDB *sqlx.DB,
	session *sessions.Session,
	discordSession *discordgo.Session,
) *http.ServeMux {
	// create a *service.TODOService type variable using the *sql.DB type variable
	var indexService = service.NewIndexService(indexDB, session, discordSession)

	// register routes
	mux := http.NewServeMux()
	middleChain := alice.New(middleware.CORS)
	mux.Handle("/", middleChain.Then(serverHandler.NewIndexHandler(indexService)))
	//mux.HandleFunc("/todos", handler.NewTODOHandler(todoService).ServeHTTP)
	return mux
}
