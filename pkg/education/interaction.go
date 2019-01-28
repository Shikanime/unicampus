package education

type Saved interface {
	School() *School
	Student() *Student
}

type Interested interface {
	School() *School
	Student() *Student
}
