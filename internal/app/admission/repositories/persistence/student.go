package persistence

import (
	"github.com/Shikanime/unicampus/pkg/admission"
	"github.com/jinzhu/gorm"
)

type Student struct {
	gorm.Model
	UUID      string
	FirstName string
	LastName  string
}

func newStudentPersistenceToDomain(d *Student) *admission.Student {
	return &admission.Student{
		UUID:      d.UUID,
		FirstName: d.FirstName,
		LastName:  d.LastName,
	}
}
