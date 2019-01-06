package persistence

import (
	"github.com/jinzhu/gorm"
)

type Repo struct {
	Conn *gorm.DB
}

func (r *Repo) GetSchool(id string) School {
	var schoolData School
	r.Conn.Where(id).Take(&schoolData)
	return schoolData
}

func (r *Repo) ListSchool(idxs []string) []School {
	schoolDatas := make([]School, len(idxs))
	if err := r.Conn.Where(idxs).Find(&schoolDatas).Error; err != nil {
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
	if err := r.Conn.Find(&schoolDatas).Error; err != nil {
		panic(err)
	}

	return schoolDatas
}
