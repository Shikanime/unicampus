package admission

type StudentPersistence interface {
	CreateStudent(student *Student) error
	GetStudent(student *Student) (*Student, error)
	UpdateStudent(student *Student) error
	DeleteStudent(student *Student) error
}

type SchoolPersistence interface {
	CreateSchool(school *School) error
	GetSchool(school *School) (*School, error)
	ListSchools(school []*School) ([]*School, error)
	UpdateSchool(school *School) error
	DeleteSchool(school *School) error
}

type ApplicationPersistence interface {
	CreateApplication(application *Application) error
}
