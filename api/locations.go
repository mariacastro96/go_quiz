package api

import (
	"encoding/json"
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
			jsonError, err := json.Marshal("There was an connection error")
			if err != nil {
				log.Println("ERROR WITH JSON MARSHAL", err)
			}
			w.WriteHeader(http.StatusServiceUnavailable)
			w.Write(jsonError)

			log.Println("QUERY ERROR", err)
		} else {
			jsonValidData, err := json.Marshal(data)
			if err != nil {
				log.Println("ERROR WITH JSON MARSHAL", err)
			}
			w.WriteHeader(http.StatusCreated)
			w.Write(jsonValidData)
		}

	}
}
