package persistence

import (
	"github.com/Shikanime/unicampus/pkg/admission"
	"github.com/jinzhu/gorm"
)

type Application struct {
	gorm.Model
	UUID        string
	studentUUID string
	schoolUUID  string
}

func newApplicationPersistenceToDomain(d *Application) *admission.Application {
	return &admission.Application{
		UUID: d.UUID,
	}
}
