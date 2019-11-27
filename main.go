package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/mariacastro96/go_quiz/api"
	"github.com/mariacastro96/go_quiz/storage"
	jsonStorage "github.com/mariacastro96/go_quiz/storage/json_storage"
	"github.com/mariacastro96/go_quiz/storage/postgres"
)

func main() {
	db, err := sql.Open("postgres", "postgres://postgres:password@localhost/quiz_locations?sslmode=disable")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()

	const path = "locations.json"
	// file, err = os.OpenFile(path, os.O_RDWR, 0644)
	// if err != nil {
	// 	log.Fatal(err)
	// 	return
	// }
	// defer file.Close()
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if err != nil {
			log.Println("lol")
			log.Fatal(err)
			return
		}
		defer file.Close()
	}

	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("lolol")
		log.Fatal(err)
		return
	}

	pgLocationsStore := postgres.LocationsRepo{
		DB: db,
	}

	fileLocationsStore := jsonStorage.LocationsRepo{
		File: file,
	}

	locationsStoreManager := storage.LocationsManager{
		PostgresRepo: pgLocationsStore,
		FileRepo:     fileLocationsStore,
	}

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/locations", api.AddLocationHandler(locationsStoreManager)).Methods("POST")
	router.HandleFunc("/locations/{id}", api.GetLocationByIDHandler(locationsStoreManager)).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
