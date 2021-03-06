package db

import (
	"database/sql"
	"strconv"

	"github.com/lucxjo/streamtasks/shared/models"
	_ "github.com/mattn/go-sqlite3"
)

func QueryAll(dbName string) []models.Task {
	db, _ := sql.Open("sqlite3", "./"+dbName)

	stmt := Prepare(db, "CREATE TABLE IF NOT EXISTS tasks (id INTEGER PRIMARY KEY AUTOINCREMENT, task TEXT, complete INTEGER)")
	stmt.Exec()
	stmt.Close()

	rowCountStmt, _ := db.Prepare("SELECT COUNT(*) FROM tasks")
	rows, _ := db.Query("SELECT id, task, complete FROM tasks")
	iterator := 0
	var rowsCount string

	rowCountStmt.QueryRow().Scan(&rowsCount)

	rowCount, _ := strconv.Atoi(rowsCount)
	var tasks []models.Task = make([]models.Task, rowCount)
	for rows.Next() {
		rows.Scan(&tasks[iterator].ID, &tasks[iterator].Task, &tasks[iterator].Complete)
		iterator++
	}

	defer db.Close()
	return tasks
}

func QueryOne(dbName string, id int) models.Task {
	db, err := sql.Open("sqlite3", "./"+dbName)

	if err != nil {
		panic(err)
	}

	stmt := Prepare(db, "CREATE TABLE IF NOT EXISTS tasks (id INTEGER PRIMARY KEY AUTOINCREMENT, task TEXT, complete INTEGER)")
	stmt.Exec()
	stmt.Close()

	rows, err := db.Query("SELECT id, task, complete FROM tasks WHERE id = ?", id)


	if err != nil {
		panic(err)
	}

	var task models.Task
	for rows.Next() {
		rows.Scan(&task.ID, &task.Task, &task.Complete)
	}

	defer db.Close()
	return task
}