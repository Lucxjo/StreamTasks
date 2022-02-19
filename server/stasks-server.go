package main

import (
	"flag"
	"fmt"
	"html/template"
	"net/http"

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

	fs := http.FileServer(http.Dir("public/static"))

	r.Handle("/static/", http.StripPrefix("/static/", fs))

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		index.Execute(w, nil)
	})

	r.HandleFunc("/add-task", func(w http.ResponseWriter, r *http.Request) {
		db.WriteDB(dbName, r.PostFormValue("task"), false)
	}).Methods("POST")

	fmt.Printf("Listening on port %s\n", port)

	http.ListenAndServe(":" + port, r)
}
