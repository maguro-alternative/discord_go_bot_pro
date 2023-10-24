package envconfig

import (
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	TOKEN            string
	DatabaseType     string
	DatabaseURL      string
	DatabaseName     string
	DatabaseUser     string
	DatabasePassword string
	DatabaseHost     string
	DatabasePort     string
	ServerPort       string
	SessionsSecret   string
	DiscordClientID  string
	DiscordSecret    string
	FrontUrl         string
	ServerUrl        string
}

func NewEnv() (*Env, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}

	return &Env{
		TOKEN:            os.Getenv("D_TOKEN"),
		DatabaseType:     "postgresql",
		DatabaseURL:      os.Getenv("PGURL"),
		DatabaseName:     os.Getenv("PGDATABASE"),
		DatabaseUser:     os.Getenv("PGUSER"),
		DatabasePassword: os.Getenv("PGPASSWORD"),
		DatabaseHost:     os.Getenv("PGHOST"),
		DatabasePort:     os.Getenv("PGPORT"),
		ServerPort:       os.Getenv("PORT"),
		SessionsSecret:   os.Getenv("SESSIONS_SECRET"),
		DiscordClientID:  os.Getenv("DISCORD_CLIENT_ID"),
		DiscordSecret:    os.Getenv("DISCORD_SECRET"),
		FrontUrl:         os.Getenv("FRONT_URL"),
		ServerUrl:        os.Getenv("SERVER_URL"),
	}, nil
}
