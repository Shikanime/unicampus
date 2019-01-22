package persistence

import (
	"github.com/jinzhu/gorm"
)

type School struct {
	gorm.Model
	UUID        string
	Name        string
	Description string
}
