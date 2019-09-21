package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/mariacastro96/go_quiz/api"
	"github.com/mariacastro96/go_quiz/locations"
)

func main() {
	db, err := sql.Open("postgres", "postgres://postgres:password@localhost/quiz_locations?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	var loc []locations.Location
	// err = db.Ping()
	// if err != nil {
	// 	log.Println("this is not ok with the DB")
	// 	// log.Fatal("Error: Could not establish a connection with the database. err: %d", err.Error)
	// }
	defer db.Close()

	log.Println("go")
	fmt.Println("db ok")

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/locations", api.AddLocationHandler(db, loc)).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}
