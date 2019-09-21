package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/mariacastro96/go_quiz/locations"
)

func AddLocationHandler(db *sql.DB, loc []locations.Location) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)

		var data locations.Location

		err := decoder.Decode(&data)
		if err != nil {
			panic(err)
		}
		log.Println("Before anything. locations array: ", loc)

		err = db.Ping()
		if err != nil {
			fmt.Fprintf(w, "Having DB problems, will save your data as soon as possible")
			loc = append(loc, data)
			log.Println("This is not working. locations array: ", loc)
		} else {
			log.Println("locations before adding to db: ", loc)
			for _, l := range loc {
				lastInsertID := 0
				log.Println("adding array locs to db: ", loc)
				err = db.QueryRow("INSERT INTO locations (lat, lon, driver_id) VALUES ($1, $2, $3) RETURNING id", l.Lat, l.Lon, l.DriverID).Scan(&lastInsertID)
				if err != nil {
					log.Fatal("QUERY ERROR", err)

				} else {
					// log.Printf("lat: %v, lon: %v, driver id: %v, id: %v", l.Lat, l.Lon, l.DriverID, lastInsertID)
					fmt.Fprintf(w, "OK! location id: %v", lastInsertID)
				}
			}
			if len(loc) > 0 {
				loc = nil
			}

			lastInsertID := 0
			err = db.QueryRow("INSERT INTO locations (lat, lon, driver_id) VALUES ($1, $2, $3) RETURNING id", data.Lat, data.Lon, data.DriverID).Scan(&lastInsertID)
			if err != nil {
				log.Fatal("QUERY ERROR", err)

			} else {
				// log.Println("Locations: ", loc)
				// log.Printf("lat: %v, lon: %v, driver id: %v, id: %v", data.Lat, data.Lon, data.DriverID, lastInsertID)

				fmt.Fprintf(w, "OK! location id: %v", lastInsertID)
			}
		}
	}
}
