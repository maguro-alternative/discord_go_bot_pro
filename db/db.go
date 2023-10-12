package db

import (
	"context"
	_ "embed"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

//go:embed schema.sql
var schema string // schema.sqlの内容をschemaに代入

var db *sqlx.DB // DBは*sql.DB型の変数、グローバル変数

type DBHandler struct {
	Driver      *sqlx.DB
	DBPing      func(context.Context) error
	CheckTables func(context.Context) (any, error)
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
	TablesCheck := func(ctx context.Context) (any, error) {
		results, err := db.ExecContext(ctx, "select schemaname, tablename, tableowner from pg_tables;")
		if err != nil {
			return nil, err
		}
		return results, nil
	}

	return &DBHandler{
		Driver:      db,
		DBPing:      PingDB,
		CheckTables: TablesCheck,
	}
}
