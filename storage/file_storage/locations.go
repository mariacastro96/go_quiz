package filestorage

import (
	"encoding/json"

	"github.com/dgraph-io/badger"
	"github.com/mariacastro96/go_quiz/locations"
)

// LocationsRepo accesses the db
type LocationsRepo struct {
	DB *badger.DB
}

// Insert locations into db
func (db LocationsRepo) Insert(data locations.Location) error {
	err := db.DB.Update(func(txn *badger.Txn) error {
		jsonData, marshalErr := json.Marshal(data)
		if marshalErr != nil {
			return marshalErr
		}
		if inserErr := txn.Set([]byte(data.ID.String()), jsonData); inserErr != nil {
			return inserErr
		}
		return nil
	})

	if err != nil {
		return err
	}
	return nil
}

// GetByID locations from db with the id
func (db LocationsRepo) GetByID(id string) (locations.Location, error) {
	var data locations.Location
	// row := pg.DB.QueryRow("SELECT id, lat, lon, driver_id FROM locations WHERE id::text=($1)", id)
	// if err := row.Scan(&data.ID, &data.Lat, &data.Lon, &data.DriverID); err != nil {
	// 	return data, err
	// }
	err := db.DB.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(id))
		if err != nil {
			return err
		}

		findErr := item.Value(func(val []byte) error {
			if unmarshalErr := json.Unmarshal(val, &data); unmarshalErr != nil {
				return unmarshalErr
			}
			return nil
		})
		if findErr != nil {
			return findErr
		}
		return nil
	})
	if err != nil {
		return data, err
	}
	return data, nil
}
