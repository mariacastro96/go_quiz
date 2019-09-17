package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"./goquiz/locations"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var data locations.location
	err := decoder.Decode(&data)
	if err != nil {
		panic(err)
	}
	owner := data.lat
	name := data.lon
	log.Println(owner, name)
	// fmt.Fprintf(w, "Welcome home!")
}

func main() {
	db, err := sql.Open("postgres", "postgres://postgres:password@localhost/quiz_locations?sslmode=disable")
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
