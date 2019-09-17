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

// AddLocationHandler adds location
// type mc struct {
// 	name string
// 	age  int
// }

func AddLocationHandler(db *sql.DB) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)

		var data locations.Location

		err := decoder.Decode(&data)
		if err != nil {
			panic(err)
		}

		locations := make(chan locations.Location, 1000)

		log.Printf("locations %T", locations)
		// go log.Println(maria)

		err = db.Ping()
		if err != nil {
			fmt.Fprintf(w, "Having DB problems, will save your data as soon as possible")
			select {
			case locations <- data:
				log.Println("received location", data)
				log.Println("LOCATIONS", len(locations))
			default:
				log.Println("no location received")
			}

		} else {
			log.Println("LOCATIONS", len(locations))
			lastInsertId := 0
			err = db.QueryRow("INSERT INTO locations (lat, lon, driver_id) VALUES ($1, $2, $3) RETURNING id", data.Lat, data.Lon, data.DriverID).Scan(&lastInsertId)
			if err != nil {
				log.Fatal("QUERY ERROR", err)

			} else {
				log.Printf("lat: %v, lon: %v, driver id: %v, id: %v", data.Lat, data.Lon, data.DriverID, lastInsertId)

				fmt.Fprintf(w, "OK! location id: %v", lastInsertId)
			}
		}
	}
}
