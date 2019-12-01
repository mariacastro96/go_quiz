package main

import (
	"database/sql"
	"log"
	"net/http"

	filestorage "github.com/mariacastro96/go_quiz/storage/file_storage"

	badger "github.com/dgraph-io/badger"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/mariacastro96/go_quiz/api"
	"github.com/mariacastro96/go_quiz/storage"
	"github.com/mariacastro96/go_quiz/storage/postgres"
)

func main() {
	db, err := sql.Open("postgres", "postgres://postgres:password@localhost/quiz_locations?sslmode=disable")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()

	fileDB, err := badger.Open(badger.DefaultOptions("file_locations"))
	if err != nil {
		log.Println("IT WAS HERE actually")
		log.Fatal(err)
		return
	}
	defer db.Close()

	const Path = "locations.json"

	pgLocationsStore := postgres.LocationsRepo{
		DB: db,
	}

	fileLocationsStore := filestorage.LocationsRepo{
		DB: fileDB,
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
