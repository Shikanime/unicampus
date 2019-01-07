package persistence

import (
	"github.com/Shikanime/unicampus/pkg/admission"
	"github.com/jinzhu/gorm"
)

type School struct {
	gorm.Model
	UUID        string
	Name        string
	Description string
}

func newSchoolPersistenceToDomain(d *School) *admission.School {
	return &admission.School{
		UUID:        d.UUID,
		Name:        d.Name,
		Description: d.Description,
	}
}

func newSchoolsPersistenceToDomain(d []*School) []*admission.School {
	schools := make([]*admission.School, len(d))
	for _, schoolData := range d {
		schools = append(schools, newSchoolPersistenceToDomain(schoolData))
	}
	return schools
}
