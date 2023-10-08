package router

import (
	"database/sql"
	"net/http"


	"github.com/maguro-alternative/discord_go_bot/server_handler"
	"github.com/maguro-alternative/discord_go_bot/service"

	"github.com/bwmarrin/discordgo"
)

func NewRouter(indexDB *sql.DB, discordSession *discordgo.Session) *http.ServeMux {
	// create a *service.TODOService type variable using the *sql.DB type variable
	var indexService = service.NewIndexService(indexDB,discordSession)

	// register routes
	mux := http.NewServeMux()
	mux.HandleFunc("/", serverHandler.NewIndexHandler(indexService).ServeHTTP)
	//mux.HandleFunc("/todos", handler.NewTODOHandler(todoService).ServeHTTP)
	return mux
}