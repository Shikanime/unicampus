package containers

import (
	"context"

	"github.com/Shikanime/unicampus/cmd/admission/services/indexer"
	"github.com/Shikanime/unicampus/cmd/admission/services/persistence"
	"github.com/Shikanime/unicampus/pkg/admission"
	"github.com/Shikanime/unicampus/pkg/unicampus_api_admission_v1"
)

func NewApplication(
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

func (s *Student) AppyStudentToSchool(ctx context.Context, in *unicampus_api_admission_v1.Application) (*unicampus_api_admission_v1.Application, error) {
	applicationData, err := s.persistence.CreateApplication(NewApplicationNetworkToDomain(in))
	if err != nil {
		return nil, err
	}

	return NewApplicationDomainToNetwork(applicationData), nil
}

func NewApplicationNetworkToDomain(application *unicampus_api_admission_v1.Application) *admission.Application {
	return &admission.Application{
		UUID:    application.Uuid,
		School:  NewSchoolNetworkToDomain(application.School),
		Student: NewStudentNetworkToDomain(application.Student),
	}
}

func NewApplicationDomainToNetwork(application *admission.Application) *unicampus_api_admission_v1.Application {
	return &unicampus_api_admission_v1.Application{
		Uuid:    application.UUID,
		School:  NewSchoolDomainToNetwork(application.School),
		Student: NewStudentDomainToNetwork(application.Student),
	}
}
