package requestdto

type LocationSearchRequest struct {
	Category string  `json:"category"`  // Category filter
	Latitude float64 `json:"latitude"`  // Latitude
	Lonitude float64 `json:"longitude"` // Longitude
	RadiusKm float64 `json:"radius_km"` // Radius in kilometers
}
