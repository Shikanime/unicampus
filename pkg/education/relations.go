package education

type Pedagogy interface {
	Name() string
	Schools() []*School
}

type Activity interface {
	Name() string
	Schools() []*School
}

type Scolarschip interface {
	Name() string
	Schools() []*School
}
