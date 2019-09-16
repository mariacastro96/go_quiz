package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	_ "github.com/lib/pq"
)

type location struct {
	ID        int     `json:"ID"`
	Latitude  float64 `json:"Title"`
	Longitude float64 `json:"Description"`
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func main() {
	db, err := sql.Open("postgres", "postgres://locations:password@localhost/mariacastro?sslmode=disable")
	if err != nil {
		fmt.Println("error:", err)
	}
	db.Exec("CREATE TABLE $1", "locations")

	fmt.Println("db ok:", db)
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/event", homeLink).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}
