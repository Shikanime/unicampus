package persistence

import (
	"github.com/jinzhu/gorm"
	"gitlab.com/deva-hub/unicampus/internal/pkg/services"
	"gitlab.com/deva-hub/unicampus/pkg/admission"
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
	if err := r.db.AutoMigrate(&School{}).Error; err != nil {
		return err
	}
	if err := r.db.AutoMigrate(&Region{}).Error; err != nil {
		return err
	}
	if err := r.db.AutoMigrate(&Student{}).Error; err != nil {
		return err
	}
	if err := r.db.AutoMigrate(&Application{}).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repo) GetSchool(school *admission.School) (*admission.School, error) {
	dbSchool := new(School)
	if err := r.db.Take(dbSchool, school).Error; err != nil {
		return nil, err
	}
	return formatSchoolDomain(*dbSchool), nil
}

func (r *Repo) ListSchools(schools []*admission.School) ([]*admission.School, error) {
	datas := make([]*School, len(schools))
	if err := r.db.Find(&datas, schools).Error; err != nil {
		return nil, err
	}

	res := make([]*admission.School, len(datas))
	for _, dbSchool := range datas {
		res = append(res, formatSchoolDomain(*dbSchool))
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
	return formatStudentDomain(*dbStudent), nil
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

func formatSchoolDomain(in School) *admission.School {
	return &admission.School{
		Identification: admission.Identification{
			UUID: in.UUID,
		},
		Name:        in.Name,
		Description: in.Description,
		Region: admission.Region{
			City:    in.Region.City,
			Country: in.Region.Country,
			State:   in.Region.State,
			Zipcode: in.Region.Zipcode,
		},
		Location: admission.Location{
			Address:   in.Address,
			Latitude:  in.Latitude,
			Longitude: in.Longitude,
		},
	}
}

func formatStudentDomain(in Student) *admission.Student {
	return &admission.Student{
		Identification: admission.Identification{
			UUID: in.UUID,
		},
		FirstName: in.FirstName,
		LastName:  in.LastName,
	}
}
