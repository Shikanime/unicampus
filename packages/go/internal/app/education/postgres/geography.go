package postgres

type Location struct {
	Address  string
	GeoPoint GeoPoint `gorm:"foreignkey:LocationRefer"`
	Region   Region   `gorm:"foreignkey:LocationRefer"`

	SchoolRefer string
}

type GeoPoint struct {
	Longitude float64
	Latitude  float64

	LocationRefer string
}

type Region struct {
	Country string
	State   string
	Zipcode string
	City    string

	LocationRefer string
}
