package app

import (
	"context"

	"gitlab.com/deva-hub/unicampus/pkg/admission"

	unicampus_api_admission_v1alpha1 "gitlab.com/deva-hub/unicampus/api/admission/v1alpha1"
	"gitlab.com/deva-hub/unicampus/cmd/admission/app/repositories/persistence"
)

func NewApplicationService(
	persistence *persistence.Repo,
) Application {
	return Application{
		persistence: persistence,
	}
}

type Application struct {
	persistence admission.ApplicationPersistence
}

func (s *Application) AppyStudentToSchool(ctx context.Context, in *unicampus_api_admission_v1alpha1.Application) (*unicampus_api_admission_v1alpha1.Application, error) {
	if err := s.persistence.CreateApplication(formatApplicationDomain(*in)); err != nil {
		return nil, err
	}

	return in, nil
}

func formatApplicationDomain(in unicampus_api_admission_v1alpha1.Application) *admission.Application {
	return &admission.Application{
		Identification: admission.Identification{
			UUID: in.UUID,
		},
		School:  formatSchoolDomain(*in.School),
		Student: formatStudentDomain(*in.Student),
	}
}

func formatApplicationNetwork(in admission.Application) *unicampus_api_admission_v1alpha1.Application {
	return &unicampus_api_admission_v1alpha1.Application{
		UUID:    in.UUID,
		School:  formatSchoolNetwork(*in.School),
		Student: formatStudentNetwork(*in.Student),
	}
}
