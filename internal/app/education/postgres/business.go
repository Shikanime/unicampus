package postgres

type School struct {
	UUID        string `gorm:"primary_key"`
	Name        string
	Description string
	Phone       string
	Email       string
	Links       []Link
	Pictures    []Link
	Locations   []Location
	Sectors     []Sector
}

type Sector struct {
	Name string
}
