package admission

type SchoolIndexer interface {
	SearchSchool(school *School) ([]*School, error)
	SearchSchoolsByQuery(query string) ([]*School, error)
	PutSchool(school *School) error
	DeleteSchool(school *School) error
}
