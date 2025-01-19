package requestdto

type TollGuruInternal struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type TollGuruRequest struct {
	From TollGuruInternal `json:"from"` // Latitude
	To   TollGuruInternal `json:"to"`   // Longitude
}
