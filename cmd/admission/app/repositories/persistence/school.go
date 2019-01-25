package persistence

type School struct {
	UUID string `gorm:"primary_key"`

	Name        string
	Description string

	Address   string
	Latitude  float64
	Longitude float64
	Region    Region `gorm:"many2many:school_regions;"`
}
