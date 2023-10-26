package envconfig

import (
	"context"
	"os"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/cockroachdb/errors"
	"github.com/joho/godotenv"
)

// RetryLimit はリトライする最大回数を指定します。
const retryLimit = 3

// RetryInterval はリトライの間隔を指定します（ミリ秒単位）。
const retryInterval = 1 // 1秒

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
	SessionsName     string
}

func NewEnv() (*Env, error) {
	EnvGet := func(ctx context.Context) error {
		operation := func() error {
			err := godotenv.Load(".env")
			return errors.WithStack(err)
		}
		return retryOperation(ctx,func() error { return operation() })
	}
	err := EnvGet(context.Background())
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
		DiscordSecret:    os.Getenv("DISCORD_CLIENT_SECRET"),
		FrontUrl:         os.Getenv("FRONT_URL"),
		ServerUrl:        os.Getenv("SERVER_URL"),
		SessionsName:     os.Getenv("SESSIONS_NAME"),
	}, nil
}

func retryOperation(ctx context.Context, operation func() error) error {
	retryBackoff := backoff.NewExponentialBackOff()
	retryBackoff.MaxElapsedTime = time.Second * retryInterval

	err := backoff.RetryNotify(func() error {
		err := operation()
		if err != nil {
			return err
		}
		err = backoff.Permanent(err)
		return errors.WithStack(err)
	}, retryBackoff, func(err error, duration time.Duration) {
		//slog.WarnContext(ctx, fmt.Sprintf("%v retrying in %v...", err, duration))
	})
	return errors.WithStack(err)
}
