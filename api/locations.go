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
func AddLocationHandler(locs storage.Postgres, st []locations.Location) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)

		var data locations.Location

		err := decoder.Decode(&data)
		if err != nil {
			panic(err)
		}

		data.ID = uuid.New()

		data, err = locs.Insert(data)
		if err != nil {
			log.Println("Sending data into channel")
			st = append(st, data)
			log.Println("QUERY ERROR", err)
			log.Println("letting the client know we good")

			fmt.Fprintf(w, "location id: %v, \nlatitude: %v, \nlongitude: %v, \ndriver id: %v", data.ID, data.Lat, data.Lon, data.DriverID)
		} else {
			fmt.Fprintf(w, "location id: %v, \nlatitude: %v, \nlongitude: %v, \ndriver id: %v", data.ID, data.Lat, data.Lon, data.DriverID)
		}
	}
}
