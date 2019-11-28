package jsonStorage

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"

	"github.com/google/uuid"
	"github.com/mariacastro96/go_quiz/locations"
)

// LocationsRepo accesses the file
type LocationsRepo struct {
	File *os.File
}

// Insert locations into file
func (file LocationsRepo) Insert(data locations.Location) error {
	fileInfo, err := file.File.Stat()
	if err != nil {
		return err
	}
	if size := fileInfo.Size(); size > 0 {
		_, writeErr := file.File.WriteString(",")
		if writeErr != nil {
			return writeErr
		}
	}
	jsonData, err := json.Marshal(data)

	if err != nil {
		return err
	}

	_, writeErr := file.File.Write(jsonData)
	if writeErr != nil {
		return writeErr
	}

	return nil
}

// GetByID locations from file with the id
func (file LocationsRepo) GetByID(id string) (locations.Location, error) {
	var data locations.Location
	var ls []locations.Location
	fileInfo, err := file.File.Stat()
	if err != nil {
		return data, err
	}
	byteValue, err := ioutil.ReadFile(fileInfo.Name())
	if err != nil {
		return data, err
	}
	locs := `[` + string(byteValue) + `]`
	json.Unmarshal([]byte(locs), &ls)
	for _, v := range ls {
		id, err := uuid.Parse(id)
		if err != nil {
			return data, err
		}
		if id == v.ID {
			return locations.Location{ID: v.ID, Lat: v.Lat, Lon: v.Lon, DriverID: v.DriverID}, nil
		}
	}
	return data, errors.New("No Rows")
}
