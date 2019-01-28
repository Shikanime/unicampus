package education

type Location struct {
	Address  string
	GeoPoint GeoPoint
	Region   Region
}

type GeoPoint struct {
	Longitude float64
	Latitude  float64
}

type Region struct {
	Country string
	State   string
	Zipcode string
	City    string
}
