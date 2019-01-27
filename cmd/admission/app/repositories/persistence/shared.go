package persistence

type Location struct {
	Address   string
	Latitude  float64
	Longitude float64
}

type Link struct {
	Name      string
	Reference string
}
