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
	Links       []Link     `gorm:"foreignkey:SchoolLink"`
	Pictures    []Link     `gorm:"foreignkey:SchoolPicture"`
	Locations   []Location `gorm:"foreignkey:SchoolLocation"`
	Sectors     []Sector   `gorm:"foreignkey:SchoolSector"`
}

type Sector struct {
	Name string
}
