package education

type SchoolStorage interface {
	CreateSchool(school *School) error
	GetSchool(school *School) (*School, error)
	ListSchools(school []*School) ([]*School, error)
	UpdateSchool(school *School) error
	DeleteSchool(school *School) error
}
