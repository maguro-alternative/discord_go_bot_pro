package db

import (
	"context"
	"database/sql"
	_ "embed"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

//go:embed schema.sql
var schema string // schema.sqlの内容をschemaに代入

var db *sqlx.DB // DBは*sql.DB型の変数、グローバル変数

type DBHandler struct {
	Driver           *sqlx.DB
	DBPing           func(context.Context) error
	CheckTables      func(context.Context) (sql.Result, error)
	QueryxContext    func(context.Context, string, ...interface{}) (*sqlx.Rows, error)
	QueryRowxContent func(context.Context, string, ...interface{}) (*sqlx.Row, error)
	GetContent       func(context.Context, interface{}, string, ...interface{}) error
	SelectContent    func(context.Context, interface{}, string, ...interface{}) error
	ExecContext      func(ctx context.Context, query string, args ...any) (*sql.Result, error)
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
	TablesCheck := func(ctx context.Context) (sql.Result, error) {
		results, err := db.ExecContext(ctx, "select schemaname, tablename, tableowner from pg_tables;")
		if err != nil {
			return nil, err
		}
		return results, nil
	}

	// QueryxContextは複数の行を返す
	QueryxContext := func(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error) {
		results, err := db.QueryxContext(ctx, query, args...)
		if err != nil {
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
		err := db.GetContext(ctx, dest, query, args...)
		if err != nil {
			return err
		}
		return nil
	}

	// SelectContentは複数の行を返す
	SelectContent := func(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
		err := db.SelectContext(ctx, dest, query, args...)
		if err != nil {
			return err
		}
		return nil
	}

	ExecContext := func(ctx context.Context, query string, args ...any) (*sql.Result, error) {
		results, err := db.ExecContext(ctx, query, args...)
		if err != nil {
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
	}
}
