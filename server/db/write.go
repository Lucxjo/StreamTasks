package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func WriteDB(dbName string, task string, complete bool) {
	db, err := sql.Open("sqlite3", "./"+dbName)

	if err != nil {
		panic(err)
	}

	stmt := Prepare(db, "CREATE TABLE IF NOT EXISTS tasks (id INTEGER PRIMARY KEY AUTOINCREMENT, task TEXT, complete INTEGER)")
	stmt.Exec()
	stmt.Close()

	insrt := Prepare(db, "INSERT INTO tasks (task, complete) VALUES (?, ?)")
	insrt.Exec(task, complete)
	insrt.Close()

	db.Close()
}