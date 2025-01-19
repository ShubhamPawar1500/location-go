package responsedto

type LocationWithDistance struct {
	ID         uint    `json:"id"`
	Name       string  `json:"name"`
	Address    string  `json:"address"`
	Category   string  `json:"category"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
	DistanceKm float64 `json:"distance"`
}
