package persistence

type Student struct {
	UUID string `gorm:"primary_key"`

	FirstName string
	LastName  string
}
