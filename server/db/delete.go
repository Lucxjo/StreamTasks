package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func DeleteAll(dbName string) {
	db, _ := sql.Open("sqlite3", "./"+dbName)
	stmt := Prepare(db, "DELETE FROM tasks")
	stmt.Exec()
	stmt.Close()
	db.Close()
}