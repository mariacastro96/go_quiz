package storage

import (
	"database/sql"

	"github.com/mariacastro96/go_quiz/locations"
)

// Postgres accesses the db
type Postgres struct {
	DB *sql.DB
}

// Insert locations into db
func (pg Postgres) Insert(data locations.Location) (locations.Location, error) {
	err := pg.DB.QueryRow("INSERT INTO locations (id, lat, lon, driver_id) VALUES ($1, $2, $3, $4) RETURNING id", data.ID, data.Lat, data.Lon, data.DriverID).Scan(&data.ID)
	if err != nil {
		return data, err
	}
	return data, nil
}
