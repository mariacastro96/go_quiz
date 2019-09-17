package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/mariacastro96/go_quiz/locations"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var data locations.Location
	err := decoder.Decode(&data)
	if err != nil {
		panic(err)
	}

	log.Printf("lat: %v, lon: %v, driver id: %v", data.Lat, data.Lon, data.DriverID)
	fmt.Fprintf(w, "lat: %v, lon: %v, driver id: %v", data.Lat, data.Lon, data.DriverID)
}

func main() {
	db, err := sql.Open("postgres", "postgres://postgres:password@localhost/quizlocations?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Error: Could not establish a connection with the database")
	}
	defer db.Close()

	log.Println("go")
	fmt.Println("db ok")

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/locations", homeLink).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}
