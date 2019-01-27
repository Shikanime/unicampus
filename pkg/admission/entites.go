package admission

type School struct {
	Identification

	Name        string
	Description string
	Location    *Location
	Region      *Region
	Pictures    []*Link
	References  []*Link
}

type Student struct {
	Identification

	FirstName string
	LastName  string
	Location  *Location
	Region    *Region
}

type Application struct {
	Identification

	School  *School
	Student *Student
}
