package containers

import (
	"github.com/Shikanime/unicampus/cmd/admission/services/indexer"
	"github.com/Shikanime/unicampus/cmd/admission/services/persistence"
)

func NewStudent(
	persistence *persistence.Repo,
	indexer *indexer.Repo,
) Student {
	return Student{
		persistence: persistence,
		indexer:     indexer,
	}
}

type Student struct {
	persistence *persistence.Repo
	indexer     *indexer.Repo
}

// func (s *Student) ListStudents(stream unicampus_admission.AdmissionService_ListStudentsServer) error {
// 	return errors.New("unmplemented")
// }

// func (s *Student) ListStudentsByQuery(in *unicampus_admission.Query, stream unicampus_admission.AdmissionService_ListStudentsByQueryServer) error {
// 	return errors.New("unmplemented")
// }

// func (s *Student) ListStudentsByOffset(in *unicampus_admission.Offset, stream unicampus_admission.AdmissionService_ListStudentsByOffsetServer) error {
// 	return errors.New("unmplemented")
// }

// func (s *Student) Get(ctx context.Context, in *unicampus_admission.Student) (*unicampus_admission.Student, error) {
// 	return &unicampus_admission.Student{}, errors.New("unmplemented")
// }

// func (s *Student) PutStudent(ctx context.Context, in *unicampus_admission.Student) (*unicampus_admission.Student, error) {
// 	return &unicampus_admission.Student{}, errors.New("unmplemented")
// }

// func (s *Student) DeleteStudent(ctx context.Context, in *unicampus_admission.Student) (*unicampus_admission.Student, error) {
// 	return &unicampus_admission.Student{}, errors.New("unmplemented")
// }
