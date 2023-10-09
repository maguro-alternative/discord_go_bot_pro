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

var DB *sql.DB
var err error

type Tx struct {
	*sql.Tx
}

type Stmt struct {
	*sql.Stmt
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
	DB, err = sql.Open("postgres", path)
	if err != nil {
		return nil, err
	}

	//if _, err := DB.Exec(schema); err != nil {
		//return nil, err
	//}

	return DB, nil
}

func PingDB() error {
	if err := DB.Ping(); err != nil {
		return err
	}
	return nil
}

func TablesCheck() (sql.Result, error) {
	results, err := DB.Exec("select schemaname, tablename, tableowner from pg_tables;");
	if err != nil {
		return nil, err
	}
	return results, nil
}
