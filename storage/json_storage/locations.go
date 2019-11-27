package jsonStorage

import (
	"encoding/json"
	"os"

	"github.com/mariacastro96/go_quiz/locations"
)

// LocationsRepo accesses the file
type LocationsRepo struct {
	File *os.File
}

// Insert locations into file
func (file LocationsRepo) Insert(data locations.Location) error {
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
	// var text = make([]byte, 1024)
	// for {
	// 	_, err = file.File.Read(text)

	// 	// break if finally arrived at end of file
	// 	if err == io.EOF {
	// 		break
	// 	}

	// 	// break if error occured
	// 	if err != nil && err != io.EOF {
	// 		isError(err)
	// 		break
	// 	}
	// }
	return data, nil
}
