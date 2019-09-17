package locations

type location struct {
	ID       int     `json:"ID"`
	Lat      float64 `json:"latitude"`
	Lon      float64 `json:"longitude"`
	DriverId float64 `json:"driver_id"`
}
