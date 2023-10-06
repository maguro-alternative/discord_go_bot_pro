package db

import (
	"database/sql"
	_ "embed"

	//"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	_ "github.com/lib/pq"
)

//go:embed schema.sql
var schema string	// schema.sqlの内容をschemaに代入

type DB struct {
	*sql.DB
}

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
	db, err := sql.Open("postgres", path)
	if err != nil {
		return nil, err
	}

	if _, err := db.Exec(schema); err != nil {
		return nil, err
	}

	return db, nil
}

// Open opens a database specified by its database driver name and a driver-specific data source name,
func (db *DB) Close() error {
	return db.DB.Close()
}

// トランザクションを開始する。
func (db *DB) Begin() (*Tx, error) {
	tx, err := db.DB.Begin()
	if err != nil {
		return nil, err
	}

	return &Tx{tx}, nil
}

// 後のクエリーや実行のために準備されたステートメントを作成します。
func (tx *Tx) Commit() error {
	return tx.Tx.Commit()
}

//
func (tx *Tx) Rollback() error {
	return tx.Tx.Rollback()
}

func (tx *Tx) Stmt(stmt *Stmt) *Stmt {
	return &Stmt{tx.Tx.Stmt(stmt.Stmt)}
}

func (stmt *Stmt) Close() error {
	return stmt.Stmt.Close()
}

func (stmt *Stmt) Exec(args ...interface{}) (sql.Result, error) {
	return stmt.Stmt.Exec(args...)
}

