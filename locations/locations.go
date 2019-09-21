package locations

import "github.com/google/uuid"

// Location bla bla bla
type Location struct {
	ID       uuid.UUID `json:"ID"`
	Lat      float64   `json:"latitude"`
	Lon      float64   `json:"longitude"`
	DriverID float64   `json:"driver_id"`
}
