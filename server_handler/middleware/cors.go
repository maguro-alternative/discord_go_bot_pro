package middleware

import (
	"net/http"

	"github.com/maguro-alternative/discord_go_bot/model/envconfig"
)

// セッションの確認
func CORS(next http.Handler) http.Handler {
	env, err := envconfig.NewEnv()
	if err != nil {
		panic(err)
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// クロスオリジン用にセット
		w.Header().Set("Access-Control-Allow-Origin", env.FrontUrl)
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "GET,PUT,POST,DELETE,UPDATE,OPTIONS")
		w.Header().Set("Content-Type", "application/json")

		// preflight用に200でいったん返す
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}
