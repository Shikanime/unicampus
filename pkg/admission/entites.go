package admission

type Location struct {
	Adresss string
	State   string
	Country string
	Code    string
}

type Link struct {
	Name      string
	Reference string
}

type School struct {
	Location
	UUID        string
	Name        string
	Description string
	Pictures    Link
	Externals   Link
}

type Student struct {
	Location
	UUID      string
	FirstName string
	LastName  string
}

type Application struct {
	UUID    string
	School  *School
	Student *Student
}
