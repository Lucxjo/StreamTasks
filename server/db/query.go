package db

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/lucxjo/streamtasks/shared/models"
	_ "github.com/mattn/go-sqlite3"
)

func QueryAll(dbName string) []models.Task {
	db, _ := sql.Open("sqlite3", "./"+dbName)
	rowCountStmt, _ := db.Prepare("SELECT COUNT(*) FROM tasks")
	rows, _ := db.Query("SELECT id, task, complete FROM tasks")
	iterator := 0
	var rowsCount string

	rowCountStmt.QueryRow().Scan(&rowsCount)

	rowCount, _ := strconv.Atoi(rowsCount)
	var tasks []models.Task = make([]models.Task, rowCount)
	for rows.Next() {
		rows.Scan(&tasks[iterator].ID, &tasks[iterator].Task, &tasks[iterator].Complete)
		fmt.Println(strconv.Itoa(tasks[iterator].ID) + ": " + tasks[iterator].Task + " " + strconv.FormatBool(tasks[iterator].Complete))
		iterator++
	}

	defer db.Close()
	return tasks
}