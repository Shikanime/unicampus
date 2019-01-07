package admission

type School struct {
	UUID        string
	Name        string
	Description string
}

type Student struct {
	UUID      string
	FirstName string
	LastName  string
}

type Application struct {
	UUID    string
	School  *School
	Student *Student
}
