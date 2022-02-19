package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func QueryAll(dbName string) {
	db, _ := sql.Open("sqlite3", "./"+dbName)
	rows, _ := db.Query("SELECT id FROM tasks")
}