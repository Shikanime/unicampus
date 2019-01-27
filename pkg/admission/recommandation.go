package admission

type SchoolRecommandation interface {
	PutSchool(school *School) error
	DeleteSchool(school *School) error
}
