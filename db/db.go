package db

import (
	"database/sql"
	_ "embed"

	//"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

//go:embed schema.sql
var schema string // schema.sqlの内容をschemaに代入

var db *sql.DB	// DBは*sql.DB型の変数、グローバル変数

type DBHandler struct {
	Driver      *sql.DB
	DBPing      func() error
	CheckTables func() (sql.Result, error)
}

// NewDB returns go-sqlite3 driver based *sql.DB.
func NewSqliteDB(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	if _, err := db.Exec(schema); err != nil {
		return nil, err
	}

	return db, nil
}

func NewPostgresDB(path string) (*sql.DB, error) {
	// データベースに接続
	db, err := sql.Open("postgres", path)
	if err != nil {
		return nil, err
	}

	// テーブルの作成
	//if _, err := DB.Exec(schema); err != nil {
		//return nil, err
	//}

	return db, nil
}

func NewDBHandler(db *sql.DB) *DBHandler {
	// データベースの接続を確認
	PingDB := func () error{
		if err := db.Ping(); err != nil {
			return err
		}
		return nil
	}

	// テーブル一覧の確認
	TablesCheck := func() (sql.Result, error) {
		results, err := db.Exec("select schemaname, tablename, tableowner from pg_tables;");
		if err != nil {
			return nil, err
		}
		return results, nil
	}
	return &DBHandler{
		Driver: db,
		DBPing: PingDB,
		CheckTables: TablesCheck,
	}
}
