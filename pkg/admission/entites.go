package admission

type School struct {
	Identification
	Location
	Region

	Name        string
	Description string

	Pictures   []Link
	References []Link
}

type Student struct {
	Identification
	Location
	Region

	FirstName string
	LastName  string
}

type Application struct {
	Identification

	School  *School
	Student *Student
}
