package persistence

import (
	"github.com/Shikanime/unicampus/internal/pkg/services"
	"github.com/Shikanime/unicampus/pkg/admission"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func NewRepository(service *services.PostgreSQLDatabaseService) Repo {
	db := service.Driver()

	db.AutoMigrate(&School{})
	db.Create(&School{Name: "ETNA", Description: "Desc", UUID: "yo"})

	return Repo{
		db: service,
	}
}

type Repo struct {
	db *services.PostgreSQLDatabaseService
}

func (r *Repo) GetSchool(school *admission.School) (*admission.School, error) {
	schoolData := new(School)
	if err := r.db.Get(schoolData, school); err != nil {
		return nil, err
	}
	return newSchoolPersistenceToDomain(schoolData), nil
}

func (r *Repo) ListSchools(schools []*admission.School) ([]*admission.School, error) {
	schoolDatas := make([]*School, len(schools))
	if err := r.db.Driver().Find(&schoolDatas, schools).Error; err != nil {
		return nil, err
	}
	return newSchoolsPersistenceToDomain(schoolDatas), nil
}

func (r *Repo) ListSchoolsByOffset(first uint64, offset uint64) ([]*admission.School, error) {
	length := first - offset
	if length < 0 {
		length = -length
	}

	schoolDatas := make([]*School, length)
	if err := r.db.Driver().Find(&schoolDatas).Error; err != nil {
		return nil, err
	}

	return newSchoolsPersistenceToDomain(schoolDatas), nil
}

func (r *Repo) CreateSchool(school *admission.School) (*admission.School, error) {
	if err := r.db.Create(&school); err != nil {
		return nil, err
	}
	schoolData := new(School)
	if err := r.db.Driver().First(&schoolData).Error; err != nil {
		return nil, err
	}
	return newSchoolPersistenceToDomain(schoolData), nil
}

func (r *Repo) UpdateSchool(school *admission.School) (*admission.School, error) {
	schoolData := new(School)
	if err := r.db.Update(schoolData); err != nil {
		return nil, err
	}
	return newSchoolPersistenceToDomain(schoolData), nil
}

func (r *Repo) DeleteSchool(school *admission.School) (*admission.School, error) {
	schoolData := new(School)
	if err := r.db.Driver().First(schoolData, school).Error; err != nil {
		return nil, err
	}
	if err := r.db.Delete(schoolData); err != nil {
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
	if err := r.db.Create(&applicationData); err != nil {
		return nil, err
	}

	return newApplicationPersistenceToDomain(applicationData), nil
}

func (r *Repo) GetStudent(student *admission.Student) (*admission.Student, error) {
	studentData := new(Student)
	if err := r.db.Get(studentData, student); err != nil {
		return nil, err
	}
	return newStudentPersistenceToDomain(studentData), nil
}

func (r *Repo) CreateStudent(student *admission.Student) (*admission.Student, error) {
	if err := r.db.Create(&student); err != nil {
		return nil, err
	}
	studentData := new(Student)
	if err := r.db.Driver().First(&studentData).Error; err != nil {
		return nil, err
	}
	return newStudentPersistenceToDomain(studentData), nil
}

func (r *Repo) UpdateStudent(student *admission.Student) (*admission.Student, error) {
	studentData := new(Student)
	if err := r.db.Update(studentData); err != nil {
		return nil, err
	}
	return newStudentPersistenceToDomain(studentData), nil
}

func (r *Repo) DeleteStudent(student *admission.Student) (*admission.Student, error) {
	studentData := new(Student)
	if err := r.db.Driver().First(studentData, student).Error; err != nil {
		return nil, err
	}
	if err := r.db.Delete(studentData); err != nil {
		return nil, err
	}
	return newStudentPersistenceToDomain(studentData), nil
}
