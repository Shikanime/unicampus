package persistence

import "github.com/jinzhu/gorm"

type School struct {
	gorm.Model
	ID          string
	Name        string
	Description string
}
