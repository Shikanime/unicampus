package app

import (
	"context"

	unicampus_api_admission_v1alpha1 "github.com/Shikanime/unicampus/api/admission/v1alpha1"
	"github.com/Shikanime/unicampus/internal/app/admission/repositories/indexer"
	"github.com/Shikanime/unicampus/internal/app/admission/repositories/persistence"
	"github.com/Shikanime/unicampus/pkg/admission"
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
	studentData, err := s.persistence.GetStudent(NewStudentNetworkToDomain(in))
	if err != nil {
		return nil, err
	}
	return NewStudentDomainToNetwork(studentData), nil
}

func (s *Student) RegisterStudent(ctx context.Context, in *unicampus_api_admission_v1alpha1.Student) (*unicampus_api_admission_v1alpha1.Student, error) {
	studentData, err := s.persistence.CreateStudent(NewStudentNetworkToDomain(in))
	if err != nil {
		return nil, err
	}
	return NewStudentDomainToNetwork(studentData), nil
}

func (s *Student) UnregisterStudent(ctx context.Context, in *unicampus_api_admission_v1alpha1.Student) (*unicampus_api_admission_v1alpha1.Student, error) {
	studentData, err := s.persistence.DeleteStudent(NewStudentNetworkToDomain(in))
	if err != nil {
		return nil, err
	}
	return NewStudentDomainToNetwork(studentData), nil
}

func (s *Student) UpdateStudent(ctx context.Context, in *unicampus_api_admission_v1alpha1.Student) (*unicampus_api_admission_v1alpha1.Student, error) {
	studentData, err := s.persistence.UpdateStudent(NewStudentNetworkToDomain(in))
	if err != nil {
		return nil, err
	}
	return NewStudentDomainToNetwork(studentData), nil
}

func NewStudentNetworkToDomain(student *unicampus_api_admission_v1alpha1.Student) *admission.Student {
	return &admission.Student{
		UUID:      student.UUID,
		FirstName: student.FirstName,
		LastName:  student.LastName,
	}
}

func NewStudentDomainToNetwork(student *admission.Student) *unicampus_api_admission_v1alpha1.Student {
	return &unicampus_api_admission_v1alpha1.Student{
		UUID:      student.UUID,
		FirstName: student.FirstName,
		LastName:  student.LastName,
	}
}
