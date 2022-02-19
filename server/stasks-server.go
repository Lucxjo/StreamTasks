package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/lucxjo/streamtasks/server/db"
)

var (
	port string
	dbName string
)

func handleFlags() {
	flag.StringVar(&port, "p", "8080", "port to listen on")
	flag.StringVar(&dbName, "db", "tasks.db", "database name")

	flag.Parse()
}

func main() {
	handleFlags()

	r := mux.NewRouter()
	index := template.Must(template.ParseFiles("public/index.html"))

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		index.Execute(w, nil)
	}).Methods("GET")

	r.HandleFunc("/add-task", func(w http.ResponseWriter, r *http.Request) {
		db.WriteDB(dbName, r.PostFormValue("task"), false)
	}).Methods("POST")

	r.HandleFunc("/complete-task", func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(r.PostFormValue("id"))
		db.CompleteTask(dbName, id)
	}).Methods("POST")

	r.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		tasks := db.QueryAll(dbName)
		data, _ := json.Marshal(tasks)
		fmt.Fprint(w, string(data))
	}).Methods("GET")

	r.HandleFunc("/task/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, _ := strconv.Atoi(vars["id"])
		tasks := db.QueryOne(dbName, id)
		data, _ := json.Marshal(tasks)
		fmt.Fprint(w, string(data))
	}).Methods("GET")

	r.HandleFunc("/clear-all", func(rw http.ResponseWriter, r *http.Request) {
		db.DeleteAll(dbName)
	}).Methods("POST")

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("public/"))).Methods("GET")

	fmt.Printf("Listening on port %s\n", port)

	http.ListenAndServe(":" + port, r)
}
