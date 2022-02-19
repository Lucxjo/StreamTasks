package models

type Task struct {
	ID int `json:"id"`
	Task string `json:"task"`
	Complete bool `json:"complete"`
}