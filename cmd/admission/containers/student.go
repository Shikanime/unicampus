package containers

import (
	"context"
	"errors"

	"github.com/Shikanime/unicampus/cmd/admission/services/indexer"
	"github.com/Shikanime/unicampus/cmd/admission/services/persistence"
	"github.com/Shikanime/unicampus/pkg/admission"
	"github.com/Shikanime/unicampus/pkg/unicampus_admission"
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

func (s *Student) GetStudent(ctx context.Context, in *unicampus_admission.Student) (*unicampus_admission.Student, error) {
	return &unicampus_admission.Student{}, errors.New("unmplemented")
}

func (s *Student) PutStudent(ctx context.Context, in *unicampus_admission.Student) (*unicampus_admission.Student, error) {
	return &unicampus_admission.Student{}, errors.New("unmplemented")
}

func (s *Student) DeleteStudent(ctx context.Context, in *unicampus_admission.Student) (*unicampus_admission.Student, error) {
	return &unicampus_admission.Student{}, errors.New("unmplemented")
}

func NewStudentNetworkToDomain(student *unicampus_admission.Student) *admission.Student {
	return &admission.Student{
		UUID:      student.Uuid,
		FirstName: student.FirstName,
		LastName:  student.LastName,
	}
}

func NewStudentDomainToNetwork(student *admission.Student) *unicampus_admission.Student {
	return &unicampus_admission.Student{
		Uuid:      student.UUID,
		FirstName: student.FirstName,
		LastName:  student.LastName,
	}
}
