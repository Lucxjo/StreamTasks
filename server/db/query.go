package db

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"github.com/lucxjo/streamtasks/shared/models"
	_ "github.com/mattn/go-sqlite3"
)

func QueryAll(dbName string) []models.Task {
	db, _ := sql.Open("sqlite3", "./"+dbName)
	rows, _ := db.Query("SELECT id, task, complete FROM tasks")
	iterator := 0
	rowCount := 0

	for rows.Next() {
		rowCount++
		fmt.Println("Row: " + strconv.Itoa(rowCount))
	}

	time.Sleep(time.Second * 1)

	var tasks []models.Task = make([]models.Task, rowCount)
	for rows.Next() {
		rows.Scan(&tasks[iterator].ID, &tasks[iterator].Task, &tasks[iterator].Complete)
		fmt.Println(strconv.Itoa(tasks[iterator].ID) + ": " + tasks[iterator].Task + " " + strconv.FormatBool(tasks[iterator].Complete))
		iterator++
	}

	defer db.Close()
	return tasks
}