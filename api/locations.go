package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/mariacastro96/go_quiz/locations"
	"github.com/mariacastro96/go_quiz/storage"
)

// AddLocationHandler decodes the json sent by client and answers to the client
func AddLocationHandler(store storage.Postgres) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)

		var data locations.Location

		err := decoder.Decode(&data)
		if err != nil {
			log.Println("DECODE ERROR", err)
		}

		data.ID = uuid.New()

		data, err = store.Insert(data)
		if err != nil {
			log.Println("QUERY ERROR", err)

			fmt.Fprintf(w, "location id: %v, \nlatitude: %v, \nlongitude: %v, \ndriver id: %v", data.ID, data.Lat, data.Lon, data.DriverID)
		} else {

			fmt.Fprintf(w, "location id: %v, \nlatitude: %v, \nlongitude: %v, \ndriver id: %v", data.ID, data.Lat, data.Lon, data.DriverID)
		}

	}
}
