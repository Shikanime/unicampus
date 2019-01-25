package persistence

type Application struct {
	UUID string `gorm:"primary_key"`

	Student     Student `gorm:"foreignkey:StudentUUID;association_foreignkey:Refer"`
	StudentUUID string

	School     School `gorm:"foreignkey:SchoolUUID;association_foreignkey:Refer"`
	SchoolUUID string
}
