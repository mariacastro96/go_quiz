package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/mariacastro96/go_quiz/locations"
	"github.com/mariacastro96/go_quiz/postgres"
)

// AddLocationHandler decodes the json sent by client and answers to the client
func AddLocationHandler(locationsRepo postgres.LocationsRepo) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)

		var data locations.Location
		if err := decoder.Decode(&data); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		data.ID = uuid.New()
		if err := locationsRepo.Insert(data); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		jsonValidData, err := json.Marshal(data)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		header := w.Header()
		header.Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(jsonValidData)
	}
}

// GetLocationByIdHandler decodes the json sent by client and answers to the client
func GetLocationByIDHandler(locationsRepo postgres.LocationsRepo) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		log.Println(params["id"])
		return
	}
}
