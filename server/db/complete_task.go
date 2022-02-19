package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func CompleteTask(dbName string, id int) {
	db, _ := sql.Open("sqlite3", "./"+dbName)
	stmt := Prepare(db, "UPDATE tasks SET complete = 1 WHERE id = ?")

	stmt.Exec(id)
}