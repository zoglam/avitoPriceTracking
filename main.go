package main

import (
	"log"
	"net/http"
	"os"

	"./dbmanager"
	_ "github.com/mattn/go-sqlite3"
)

func handlerInit(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func handlerSaveData(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	url := r.PostForm.Get("url")
	email := r.PostForm.Get("email")

	db := dbmanager.OpenDbConnection("sqlite3.db")
	defer dbmanager.CloseDbConnection(db)

	dbmanager.AddUrl(db, url, 0)
	urlID := dbmanager.GetUrlID(db, url)
	dbmanager.AddSubscription(db, email, urlID)

	http.Redirect(w, r, "/", 302)
}

func main() {
	args := os.Args[1:]
	for i := 0; i < len(args); i++ {
		if args[i] == "--reset" {
			dbmanager.DbReset()
		}
	}

	http.HandleFunc("/", handlerInit)
	http.HandleFunc("/save/", handlerSaveData)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
