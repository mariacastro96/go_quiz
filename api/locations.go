package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/mariacastro96/go_quiz/locations"
	"github.com/mariacastro96/go_quiz/postgres"
)


func AddLocationHandler(db *sql.DB) func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)

		var data locations.Location

		err := decoder.Decode(&data)
		if err != nil {
			panic(err)
		}

		data.ID = uuid.New()

		err = postgres.InsertLocations(data)
		if err != nil {
			log.Println("QUERY ERROR", err)
		} else {
			fmt.Fprintf(w, "location id: %v, \nlatitude: %v, \nlongitude: %v, \ndriver id: %v", data.ID, data.Lat, data.Lon, data.DriverID)
		}
	}
}
