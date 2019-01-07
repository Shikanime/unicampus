package persistence

import (
	"log"

	"github.com/Shikanime/unicampus/cmd/school/domain"
	"github.com/jinzhu/gorm"
)

func NewClient() *Repo {
	conn, err := gorm.Open("sqlite3", "gorm.db")
	if err != nil {
		log.Fatalf("failed to connect persistent database: %v", err)
	}

	// Migrate
	conn.AutoMigrate(&School{})

	// Seed
	conn.Create(&School{ID: "1", Name: "ETNA", Description: "Desc"})

	return &Repo{
		conn: conn,
	}
}

type Repo struct {
	conn *gorm.DB
}

func (r *Repo) GetSchool(school *domain.School) *domain.School {
	var schoolData *School
	if err := r.conn.First(schoolData, school).Error; err != nil {
		panic(err)
	}
	return newSchoolPersistenceToDomain(schoolData)
}

func (r *Repo) ListSchool(schools []*domain.School) []*domain.School {
	schoolDatas := make([]*School, len(schools))
	if err := r.conn.Find(&schoolDatas, schools).Error; err != nil {
		panic(err)
	}
	return newSchoolsPersistenceToDomain(schoolDatas)
}

func (r *Repo) ListSchoolByOffset(first uint64, offset uint64) []*domain.School {
	length := first - offset
	if length < 0 {
		length = -length
	}

	schoolDatas := make([]*School, length)
	if err := r.conn.Find(&schoolDatas).Error; err != nil {
		panic(err)
	}

	return newSchoolsPersistenceToDomain(schoolDatas)
}

func (r *Repo) CreateSchool(school *domain.School) *domain.School {
	if err := r.conn.Create(&school).Error; err != nil {
		panic(err)
	}
	var schoolData *School
	if err := r.conn.First(&schoolData).Error; err != nil {
		panic(err)
	}
	return newSchoolPersistenceToDomain(schoolData)
}

func (r *Repo) UpdateSchool(school *domain.School) *domain.School {
	var schoolData *School
	if err := r.conn.Update(schoolData).Error; err != nil {
		panic(err)
	}
	return newSchoolPersistenceToDomain(schoolData)
}

func (r *Repo) PutSchool(school *domain.School) *domain.School {
	var schoolData *domain.School
	if r.conn.NewRecord(school) {
		schoolData = r.CreateSchool(school)
	} else {
		schoolData = r.UpdateSchool(school)
	}
	return schoolData
}

func (r *Repo) DeleteSchool(school *domain.School) *domain.School {
	var schoolData *School
	if err := r.conn.First(schoolData, school).Error; err != nil {
		panic(err)
	}
	if err := r.conn.Delete(schoolData).Error; err != nil {
		panic(err)
	}
	return newSchoolPersistenceToDomain(schoolData)
}

func (r *Repo) Close() {
	r.conn.Close()
}
