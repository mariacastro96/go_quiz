package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/mariacastro96/go_quiz/api"
	"github.com/mariacastro96/go_quiz/postgres"
)

func main() {
	db, err := sql.Open("postgres", "postgres://postgres:password@localhost/quiz_locations?sslmode=disable")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()

	locationsStore := postgres.LocationsRepo{
		DB: db,
	}

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/locations", api.AddLocationHandler(locationsStore)).Methods("POST")
	router.HandleFunc("/locations/{id}", api.GetLocationByIDHandler(locationsStore)).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
