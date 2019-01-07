package persistence

import (
	"github.com/Shikanime/unicampus/cmd/school/domain"
	"github.com/jinzhu/gorm"
)

type School struct {
	gorm.Model
	ID          string
	Name        string
	Description string
}

func newSchoolPersistenceToDomain(d *School) *domain.School {
	return &domain.School{
		ID:          d.ID,
		Name:        d.Name,
		Description: d.Description,
	}
}

func newSchoolsPersistenceToDomain(d []*School) []*domain.School {
	schools := make([]*domain.School, len(d))
	for _, schoolData := range d {
		schools = append(schools, newSchoolPersistenceToDomain(schoolData))
	}
	return schools
}
