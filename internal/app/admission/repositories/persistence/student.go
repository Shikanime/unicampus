package persistence

import (
	"github.com/jinzhu/gorm"
)

type Student struct {
	gorm.Model
	UUID      string
	FirstName string
	LastName  string
}
