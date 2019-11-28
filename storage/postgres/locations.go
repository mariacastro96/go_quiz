package postgres

import (
	"database/sql"

	"github.com/mariacastro96/go_quiz/locations"
)

// LocationsRepo accesses the db
type LocationsRepo struct {
	DB *sql.DB
}

// Insert locations into db
func (pg LocationsRepo) Insert(data locations.Location) error {
	if _, err := pg.DB.Exec("INSERT INTO locations (id, lat, lon, driver_id) VALUES ($1, $2, $3, $4)", data.ID, data.Lat, data.Lon, data.DriverID); err != nil {
		return err
	}
	return nil
}

// GetByID locations from db with the id
func (pg LocationsRepo) GetByID(id string) (locations.Location, error) {
	var data locations.Location
	row := pg.DB.QueryRow("SELECT id, lat, lon, driver_id FROM locations WHERE id::text=($1)", id)
	if err := row.Scan(&data.ID, &data.Lat, &data.Lon, &data.DriverID); err != nil {
		return data, err
	}

	return data, nil
}
