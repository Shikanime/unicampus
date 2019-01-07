package persistence

import (
	"log"

	"github.com/Shikanime/unicampus/pkg/admission"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func NewClient() Repo {
	conn, err := gorm.Open("postgres", "sslmode=disable user=postgres password=postgres dbname=yo")
	if err != nil {
		log.Fatalf("failed to connect persistent database: %v", err)
	}

	// Migrate
	if err != nil {
		log.Fatalf("failed to connect persistent database: %v", err)
	}

	conn.AutoMigrate(&School{})

	// Seed
	conn.Create(&School{Name: "ETNA", Description: "Desc", UUID: "yo"})

	return Repo{
		conn: conn,
	}
}

type Repo struct {
	conn *gorm.DB
}

func (r *Repo) GetSchool(school *admission.School) (*admission.School, error) {
	schoolData := new(School)
	if err := r.conn.Take(schoolData, school).Error; err != nil {
		return nil, err
	}
	return newSchoolPersistenceToDomain(schoolData), nil
}

func (r *Repo) ListSchools(schools []*admission.School) ([]*admission.School, error) {
	schoolDatas := make([]*School, len(schools))
	if err := r.conn.Find(&schoolDatas, schools).Error; err != nil {
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
	if err := r.conn.Find(&schoolDatas).Error; err != nil {
		return nil, err
	}

	return newSchoolsPersistenceToDomain(schoolDatas), nil
}

func (r *Repo) CreateSchool(school *admission.School) (*admission.School, error) {
	if err := r.conn.Create(&school).Error; err != nil {
		return nil, err
	}
	schoolData := new(School)
	if err := r.conn.First(&schoolData).Error; err != nil {
		return nil, err
	}
	return newSchoolPersistenceToDomain(schoolData), nil
}

func (r *Repo) UpdateSchool(school *admission.School) (*admission.School, error) {
	schoolData := new(School)
	if err := r.conn.Update(schoolData).Error; err != nil {
		return nil, err
	}
	return newSchoolPersistenceToDomain(schoolData), nil
}

func (r *Repo) PutSchool(school *admission.School) (*admission.School, error) {
	if r.conn.NewRecord(school) {
		return r.CreateSchool(school)
	}
	return r.UpdateSchool(school)
}

func (r *Repo) DeleteSchool(school *admission.School) (*admission.School, error) {
	schoolData := new(School)
	if err := r.conn.First(schoolData, school).Error; err != nil {
		return nil, err
	}
	if err := r.conn.Delete(schoolData).Error; err != nil {
		return nil, err
	}
	return newSchoolPersistenceToDomain(schoolData), nil
}

func (r *Repo) GetStudent(student *admission.Student) (*admission.Student, error) {
	return nil, nil
}

func (r *Repo) ListStudent(schools []*admission.Student) ([]*admission.Student, error) {
	return nil, nil
}

func (r *Repo) ListStudentByOffset(first uint64, offset uint64) ([]*admission.Student, error) {
	return nil, nil
}

func (r *Repo) PutStudent(student *admission.Student) (*admission.Student, error) {
	return nil, nil
}

func (r *Repo) DeleteStudent(student *admission.Student) (*admission.Student, error) {
	return nil, nil
}

func (r *Repo) Close() {
	r.conn.Close()
}
