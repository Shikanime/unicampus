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

func (r *Repo) CreateSchool(school *admission.School) (*admission.School, error) {
	if err := r.db.Create(&school).Error; err != nil {
		return nil, err
	}
	schoolData := new(School)
	if err := r.db.First(&schoolData).Error; err != nil {
		return nil, err
	}
	return newSchoolPersistenceToDomain(schoolData), nil
}

func (r *Repo) UpdateSchool(school *admission.School) (*admission.School, error) {
	schoolData := new(School)
	if err := r.db.Update(schoolData).Error; err != nil {
		return nil, err
	}
	return newSchoolPersistenceToDomain(schoolData), nil
}

func (r *Repo) DeleteSchool(school *admission.School) (*admission.School, error) {
	schoolData := new(School)
	if err := r.db.First(schoolData, school).Error; err != nil {
		return nil, err
	}
	if err := r.db.Delete(schoolData).Error; err != nil {
		return nil, err
	}
	return newSchoolPersistenceToDomain(schoolData), nil
}

func (r *Repo) CreateApplication(application *admission.Application) (*admission.Application, error) {
	schoolData, err := r.GetSchool(application.School)
	if err != nil {
		return nil, err
	}

	studentData, err := r.GetStudent(application.Student)
	if err != nil {
		return nil, err
	}

	applicationData := &Application{
		SchoolUUID:  schoolData.UUID,
		StudentUUID: studentData.UUID,
	}
	if err := r.db.Create(&applicationData).Error; err != nil {
		return nil, err
	}

	return newApplicationPersistenceToDomain(applicationData), nil
}

func (r *Repo) GetStudent(student *admission.Student) (*admission.Student, error) {
	studentData := new(Student)
	if err := r.db.Take(studentData, student).Error; err != nil {
		return nil, err
	}
	return newStudentPersistenceToDomain(studentData), nil
}

func (r *Repo) CreateStudent(student *admission.Student) (*admission.Student, error) {
	if err := r.db.Create(&student).Error; err != nil {
		return nil, err
	}
	studentData := new(Student)
	if err := r.db.First(&studentData).Error; err != nil {
		return nil, err
	}
	return newStudentPersistenceToDomain(studentData), nil
}

func (r *Repo) UpdateStudent(student *admission.Student) (*admission.Student, error) {
	studentData := new(Student)
	if err := r.db.Update(studentData).Error; err != nil {
		return nil, err
	}
	return newStudentPersistenceToDomain(studentData), nil
}

func (r *Repo) DeleteStudent(student *admission.Student) (*admission.Student, error) {
	studentData := new(Student)
	if err := r.db.First(studentData, student).Error; err != nil {
		return nil, err
	}
	if err := r.db.Delete(studentData).Error; err != nil {
		return nil, err
	}
	return newStudentPersistenceToDomain(studentData), nil
}
