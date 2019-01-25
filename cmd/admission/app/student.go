package app

import (
	"context"

	"gitlab.com/deva-hub/unicampus/pkg/admission"

	unicampus_api_admission_v1alpha1 "gitlab.com/deva-hub/unicampus/api/admission/v1alpha1"
	"gitlab.com/deva-hub/unicampus/cmd/admission/app/repositories/persistence"
)

func NewStudentService(persistence *persistence.Repo) Student {
	return Student{
		persistence: persistence,
	}
}

type Student struct {
	persistence admission.StudentPersistence
}

func (s *Student) FindStudent(ctx context.Context, in *unicampus_api_admission_v1alpha1.Student) (*unicampus_api_admission_v1alpha1.Student, error) {
	student, err := s.persistence.GetStudent(formatStudentDomain(*in))
	if err != nil {
		return nil, err
	}
	return formatStudentNetwork(*student), nil
}

func (s *Student) RegisterStudent(ctx context.Context, in *unicampus_api_admission_v1alpha1.Student) (*unicampus_api_admission_v1alpha1.Student, error) {
	if err := s.persistence.CreateStudent(formatStudentDomain(*in)); err != nil {
		return nil, err
	}
	return in, nil
}

func (s *Student) UnregisterStudent(ctx context.Context, in *unicampus_api_admission_v1alpha1.Student) (*unicampus_api_admission_v1alpha1.Student, error) {
	school := formatStudentDomain(*in)
	if err := s.persistence.DeleteStudent(school); err != nil {
		return nil, err
	}
	return in, nil
}

func (s *Student) UpdateStudent(ctx context.Context, in *unicampus_api_admission_v1alpha1.Student) (*unicampus_api_admission_v1alpha1.Student, error) {
	school := formatStudentDomain(*in)
	if err := s.persistence.UpdateStudent(school); err != nil {
		return nil, err
	}
	return in, nil
}

func formatStudentDomain(in unicampus_api_admission_v1alpha1.Student) *admission.Student {
	return &admission.Student{
		Identification: admission.Identification{
			UUID: in.UUID,
		},
		FirstName: in.FirstName,
		LastName:  in.LastName,
	}
}

func formatStudentNetwork(in admission.Student) *unicampus_api_admission_v1alpha1.Student {
	return &unicampus_api_admission_v1alpha1.Student{
		UUID:      in.UUID,
		FirstName: in.FirstName,
		LastName:  in.LastName,
	}
}
