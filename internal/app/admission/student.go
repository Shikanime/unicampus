package admission

import (
	"context"

	unicampus_api_admission_v1alpha1 "github.com/Shikanime/unicampus/api/admission/v1alpha1"
	"github.com/Shikanime/unicampus/internal/app/admission/repositories/indexer"
	"github.com/Shikanime/unicampus/internal/app/admission/repositories/persistence"
	"github.com/Shikanime/unicampus/pkg/objconv"
)

func NewStudentService(
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

func (s *Student) FindStudent(ctx context.Context, in *unicampus_api_admission_v1alpha1.Student) (*unicampus_api_admission_v1alpha1.Student, error) {
	studentData, err := s.persistence.GetStudent(objconv.FormatStudentDomain(in))
	if err != nil {
		return nil, err
	}
	return objconv.FormatStudentNetwork(studentData), nil
}

func (s *Student) RegisterStudent(ctx context.Context, in *unicampus_api_admission_v1alpha1.Student) (*unicampus_api_admission_v1alpha1.Student, error) {
	if err := s.persistence.CreateStudent(objconv.FormatStudentDomain(in)); err != nil {
		return nil, err
	}
	return in, nil
}

func (s *Student) UnregisterStudent(ctx context.Context, in *unicampus_api_admission_v1alpha1.Student) (*unicampus_api_admission_v1alpha1.Student, error) {
	school := objconv.FormatStudentDomain(in)
	if err := s.persistence.DeleteStudent(school); err != nil {
		return nil, err
	}
	return in, nil
}

func (s *Student) UpdateStudent(ctx context.Context, in *unicampus_api_admission_v1alpha1.Student) (*unicampus_api_admission_v1alpha1.Student, error) {
	school := objconv.FormatStudentDomain(in)
	if err := s.persistence.UpdateStudent(school); err != nil {
		return nil, err
	}
	return in, nil
}
