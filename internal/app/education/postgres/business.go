package postgres

import "time"

type School struct {
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time `sql:"index"`
	UUID        string     `gorm:"primary_key"`
	Name        string
	Description string
	Phone       string
	Email       string
	Links       []Link     `gorm:"foreignkey:SchoolRefer"`
	Pictures    []Link     `gorm:"foreignkey:SchoolRefer"`
	Locations   []Location `gorm:"foreignkey:SchoolRefer"`
	Sectors     []Sector   `gorm:"foreignkey:SchoolRefer"`
}

type Sector struct {
	Name string
}
