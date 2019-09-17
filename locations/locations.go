package locations

// Location bla bla bla
type Location struct {
	ID       int     `json:"ID"`
	Lat      float64 `json:"latitude"`
	Lon      float64 `json:"longitude"`
	DriverID float64 `json:"driver_id"`
}
