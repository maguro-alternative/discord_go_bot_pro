package db

import (
	"context"
	"database/sql"
	_ "embed"
	"time"

	"github.com/maguro-alternative/discord_go_bot/db/table"

	"github.com/cenkalti/backoff/v4"
	"github.com/cockroachdb/errors"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

//go:embed schema.sql
var schema string // schema.sqlの内容をschemaに代入

var db *sqlx.DB // DBは*sql.DB型の変数、グローバル変数

// RetryLimit はリトライする最大回数を指定します。
const retryLimit = 3

// RetryInterval はリトライの間隔を指定します（ミリ秒単位）。
const retryInterval = 1 // 1秒


type DBHandler struct {
	Driver           *sqlx.DB
	DBPing           func(context.Context) error
	CheckTables      func(context.Context,[]table.PGTable) error
	QueryxContext    func(context.Context, string, ...interface{}) (*sqlx.Rows, error)
	QueryRowxContent func(context.Context, string, ...interface{}) (*sqlx.Row, error)
	GetContent       func(context.Context, interface{}, string, ...interface{}) error
	SelectContent    func(context.Context, interface{}, string, ...interface{}) error
	ExecContext      func(ctx context.Context, query string, args ...any) (*sql.Result, error)
	NamedExecContext func(ctx context.Context, query string, arg interface{}) (*sql.Result, error)
}

// NewDB returns go-sqlite3 driver based *sql.DB.
func NewSqliteDB(path string) (*sqlx.DB, error) {
	db, err := sqlx.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	if _, err := db.Exec(schema); err != nil {
		return nil, err
	}

	return db, nil
}

func NewPostgresDB(path string) (*sqlx.DB, error) {
	// データベースに接続
	db, err := sqlx.Open("postgres", path)
	if err != nil {
		return nil, err
	}

	// テーブルの作成
	//if _, err := DB.Exec(schema); err != nil {
	//return nil, err
	//}

	return db, nil
}

func NewDBHandler(db *sqlx.DB) *DBHandler {
	/*
		データベースで行う処理をまとめた構造体を返す

		引数
			db: *sql.DB型の変数

		戻り値
			*DBHandler型の変数
	*/
	// データベースの接続を確認
	PingDB := func(ctx context.Context) error {
		if err := db.PingContext(ctx); err != nil {
			return err
		}
		return nil
	}

	// テーブル一覧の確認
	TablesCheck := func(ctx context.Context, dest []table.PGTable) error {
		// retryOperationはエラーが発生した場合にリトライする
		operation := func() error {
			err := db.SelectContext(ctx, dest, "select schemaname, tablename, tableowner from pg_tables;")
			return errors.WithStack(err)
		}
		return retryOperation(ctx,func() error { return operation() })
	}

	// QueryxContextは複数の行を返す
	QueryxContext := func(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error) {
		var err error
		var results *sqlx.Rows
		operation := func() error {
			results, err = db.QueryxContext(ctx, query, args...)
			return errors.WithStack(err)
		}
		if err := retryOperation(
			ctx,
			func() error { return operation() },
		); err != nil {
			return nil, err
		}

		return results, nil
	}

	// QueryRowxContextは1行を返す
	QueryRowxContent := func(ctx context.Context, query string, args ...interface{}) (*sqlx.Row, error) {
		results := db.QueryRowxContext(ctx, query, args...)
		return results, nil
	}

	// GetContentは1行を返す
	GetContent := func(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
		// retryOperationはエラーが発生した場合にリトライする
		operation := func() error {
			err := db.GetContext(ctx, dest, query, args...)
			return errors.WithStack(err)
		}
		return retryOperation(
			ctx,
			func() error {
				return operation()
			},
		)
	}

	// SelectContentは複数の行を返す
	SelectContent := func(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
		// retryOperationはエラーが発生した場合にリトライする
		operation := func() error {
			err := db.SelectContext(ctx, dest, query, args...)
			return errors.WithStack(err)
		}
		return retryOperation(
			ctx,
			func() error {
				return operation()
			},
		)
	}

	// ExecContextは複数の行を返す
	ExecContext := func(ctx context.Context, query string, args ...any) (*sql.Result, error) {
		var err error
		var results sql.Result
		// retryOperationはエラーが発生した場合にリトライする
		operation := func() error {
			results, err = db.ExecContext(ctx, query, args...)
			return errors.WithStack(err)
		}
		if err := retryOperation(
			ctx,
			func() error { return operation() },
		); err != nil {
			return nil, err
		}

		return &results, nil
	}

	// NamedExecContextは複数の行を返す
	NamedExecContext := func(ctx context.Context, query string, arg interface{}) (*sql.Result, error) {
		var err error
		var results sql.Result
		// retryOperationはエラーが発生した場合にリトライする
		operation := func() error {
			results, err = db.NamedExecContext(ctx, query, arg)
			return errors.WithStack(err)
		}
		if err := retryOperation(
			ctx,
			func() error { return operation() },
		); err != nil {
			return nil, err
		}

		return &results, nil
	}

	return &DBHandler{
		Driver:           db,
		DBPing:           PingDB,
		CheckTables:      TablesCheck,
		QueryxContext:    QueryxContext,
		QueryRowxContent: QueryRowxContent,
		GetContent:       GetContent,
		SelectContent:    SelectContent,
		ExecContext:      ExecContext,
		NamedExecContext: NamedExecContext,
	}
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
