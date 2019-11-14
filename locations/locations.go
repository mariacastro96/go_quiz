package locations

import "github.com/google/uuid"

// Location bla bla bla
type Location struct {
	ID       uuid.UUID `json:"id"`
	Lat      float64   `json:"latitude"`
	Lon      float64   `json:"longitude"`
	DriverID int       `json:"driver_id"`
}
