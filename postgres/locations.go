package postgres

import (
	"database/sql"

	"github.com/mariacastro96/go_quiz/locations"
)

type postgres struct {
	DB *sql.DB
}

func (pg *postgres) InsertLocations(data locations.Location) locations.Location {
	pg.DB.QueryRow("INSERT INTO locations (id, lat, lon, driver_id) VALUES ($1, $2, $3, $4) RETURNING id lat lon driver_id", data.ID, data.Lat, data.Lon, data.DriverID).Scan(&data.ID, &data.Lat, &data.Lon, &data.DriverID)
	return data
}
