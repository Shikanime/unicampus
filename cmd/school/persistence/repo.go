package persistence

import (
	"log"

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

func (r *Repo) GetSchool(id string) School {
	var schoolData School
	r.conn.Where(id).Take(&schoolData)
	return schoolData
}

func (r *Repo) ListSchool(idxs []string) []School {
	schoolDatas := make([]School, len(idxs))
	if err := r.conn.Where(idxs).Find(&schoolDatas).Error; err != nil {
		panic(err)
	}

	return schoolDatas
}

func (r *Repo) ListSchoolByOffset(first uint64, offset uint64) []School {
	length := first - offset
	if length < 0 {
		length = -length
	}

	schoolDatas := make([]School, length)
	if err := r.conn.Find(&schoolDatas).Error; err != nil {
		panic(err)
	}

	return schoolDatas
}

func (r *Repo) Close() {
	r.conn.Close()
}
