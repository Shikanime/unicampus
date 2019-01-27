package persistence

type School struct {
	Location
	UUID        string `gorm:"primary_key"`
	Name        string
	Description string
	Region      Region `gorm:"many2many:school_regions;"`
	References  []Link
}
