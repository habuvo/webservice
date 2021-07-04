package storage

import (
	"context"
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type SQLiteStorage struct {
	db *sql.DB
}

func NewSQLiteStorage(addr string) (sst *SQLiteStorage, err error) {
	os.Remove(addr)
	file, err := os.Create(addr)
	if err != nil {
		return
	}
	file.Close()
	log.Println("sqlite DB created")

	db, err := sql.Open("sqlite3", "./"+addr)
	if err != nil {
		return
	}
	sst.db = db

	return
}

func (sq *SQLiteStorage) GetConnect() (*sql.Conn, error) {
	return sq.db.Conn(context.Background())
}
