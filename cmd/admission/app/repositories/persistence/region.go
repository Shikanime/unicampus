package persistence

type Region struct {
	UUID string `gorm:"primary_key"`

	City    string
	State   string
	Country string
	Zipcode string
}
