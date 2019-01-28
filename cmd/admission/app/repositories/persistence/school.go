package persistence

type School struct {
	UUID        string `gorm:"primary_key"`
	Name        string
	Description string
	Location    Location `gorm:"has_one:school_locations;"`
	Region      Region   `gorm:"many2many:school_regions;"`
	References  []Link
}
