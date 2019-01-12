package app

import (
	"context"

	unicampus_api_admission_v1alpha1 "github.com/Shikanime/unicampus/api/admission/v1alpha1"
	"github.com/Shikanime/unicampus/internal/app/admission/repositories/indexer"
	"github.com/Shikanime/unicampus/internal/app/admission/repositories/persistence"
	"github.com/Shikanime/unicampus/pkg/admission"
)

func NewApplicationService(
	persistence *persistence.Repo,
	indexer *indexer.Repo,
) Application {
	return Application{
		persistence: persistence,
		indexer:     indexer,
	}
}

type Application struct {
	persistence *persistence.Repo
	indexer     *indexer.Repo
}

func (s *Student) AppyStudentToSchool(ctx context.Context, in *unicampus_api_admission_v1alpha1.Application) (*unicampus_api_admission_v1alpha1.Application, error) {
	applicationData, err := s.persistence.CreateApplication(NewApplicationNetworkToDomain(in))
	if err != nil {
		return nil, err
	}

	return NewApplicationDomainToNetwork(applicationData), nil
}

func NewApplicationNetworkToDomain(application *unicampus_api_admission_v1alpha1.Application) *admission.Application {
	return &admission.Application{
		UUID:    application.Uuid,
		School:  NewSchoolNetworkToDomain(application.School),
		Student: NewStudentNetworkToDomain(application.Student),
	}
}

func NewApplicationDomainToNetwork(application *admission.Application) *unicampus_api_admission_v1alpha1.Application {
	return &unicampus_api_admission_v1alpha1.Application{
		Uuid:    application.UUID,
		School:  NewSchoolDomainToNetwork(application.School),
		Student: NewStudentDomainToNetwork(application.Student),
	}
}
