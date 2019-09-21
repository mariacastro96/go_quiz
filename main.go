package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/mariacastro96/go_quiz/api"
)

func main() {
	db, err := sql.Open("postgres", "postgres://postgres:password@localhost/quiz_locations?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/locations", api.AddLocationHandler(db)).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}
