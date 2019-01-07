package persistence

import "github.com/jinzhu/gorm"

type Student struct {
	gorm.Model
	FirstName string
	LastName  string
}
