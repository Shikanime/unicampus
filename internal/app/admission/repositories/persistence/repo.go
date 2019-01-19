package persistence

import (
	"github.com/Shikanime/unicampus/internal/pkg/services"
	"github.com/Shikanime/unicampus/pkg/admission"
	"github.com/jinzhu/gorm"
)

func NewRepository(service *services.PostgreSQLService) Repo {
	return Repo{
		db: service.Driver(),
	}
}

type Repo struct {
	db *gorm.DB
}

func (r *Repo) Init() error {
	return r.db.AutoMigrate(&School{}).Error
}

func (r *Repo) GetSchool(school *admission.School) (*admission.School, error) {
	schoolData := new(School)
	if err := r.db.Take(schoolData, school).Error; err != nil {
		return nil, err
	}
	return newSchoolPersistenceToDomain(schoolData), nil
}

func (r *Repo) ListSchools(schools []*admission.School) ([]*admission.School, error) {
	schoolDatas := make([]*School, len(schools))
	if err := r.db.Find(&schoolDatas, schools).Error; err != nil {
		return nil, err
	}
	return newSchoolsPersistenceToDomain(schoolDatas), nil
}

func (r *Repo) CreateSchool(school *admission.School) error {
	if err := r.db.Create(school).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repo) UpdateSchool(school *admission.School) error {
	if err := r.db.Update(school).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repo) DeleteSchool(school *admission.School) error {
	if err := r.db.Delete(school).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repo) CreateApplication(application *admission.Application) error {
	schoolData, err := r.GetSchool(application.School)
	if err != nil {
		return err
	}

	studentData, err := r.GetStudent(application.Student)
	if err != nil {
		return err
	}

	if err := r.db.Create(&Application{
		SchoolUUID:  schoolData.UUID,
		StudentUUID: studentData.UUID,
	}).Error; err != nil {
		return err
	}

	return nil
}

func (r *Repo) GetStudent(student *admission.Student) (*admission.Student, error) {
	studentData := new(Student)
	if err := r.db.Take(studentData, student).Error; err != nil {
		return nil, err
	}
	return newStudentPersistenceToDomain(studentData), nil
}

func (r *Repo) CreateStudent(student *admission.Student) error {
	if err := r.db.Create(student).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repo) UpdateStudent(student *admission.Student) error {
	if err := r.db.Update(student).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repo) DeleteStudent(student *admission.Student) error {
	if err := r.db.Delete(student).Error; err != nil {
		return err
	}
	return nil
}
