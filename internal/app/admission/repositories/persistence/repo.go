package persistence

import (
	"github.com/Shikanime/unicampus/internal/pkg/services"
	"github.com/Shikanime/unicampus/pkg/admission"
	"github.com/Shikanime/unicampus/pkg/objconv"
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
	dbSchool := new(School)
	if err := r.db.Take(dbSchool, school).Error; err != nil {
		return nil, err
	}
	return objconv.FormatSchoolDomain(dbSchool), nil
}

func (r *Repo) ListSchools(schools []*admission.School) ([]*admission.School, error) {
	datas := make([]*School, len(schools))
	if err := r.db.Find(&datas, schools).Error; err != nil {
		return nil, err
	}

	res := make([]*admission.School, len(datas))
	for _, dbSchool := range datas {
		res = append(res, objconv.FormatSchoolDomain(dbSchool))
	}

	return res, nil
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
	dbSchool, err := r.GetSchool(application.School)
	if err != nil {
		return err
	}

	dbStudent, err := r.GetStudent(application.Student)
	if err != nil {
		return err
	}

	if err := r.db.Create(&Application{
		SchoolUUID:  dbSchool.UUID,
		StudentUUID: dbStudent.UUID,
	}).Error; err != nil {
		return err
	}

	return nil
}

func (r *Repo) GetStudent(student *admission.Student) (*admission.Student, error) {
	dbStudent := new(Student)
	if err := r.db.Take(dbStudent, student).Error; err != nil {
		return nil, err
	}
	return objconv.FormatStudentDomain(dbStudent), nil
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
