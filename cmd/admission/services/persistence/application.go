package persistence

import (
	"github.com/Shikanime/unicampus/pkg/admission"
	"github.com/jinzhu/gorm"
)

type Application struct {
	gorm.Model
	UUID string

	Student     Student `gorm:"foreignkey:StudentUUID;association_foreignkey:Refer"`
	StudentUUID string

	School     School `gorm:"foreignkey:SchoolUUID;association_foreignkey:Refer"`
	SchoolUUID string
}

func newApplicationPersistenceToDomain(d *Application) *admission.Application {
	return &admission.Application{
		UUID: d.UUID,
	}
}
