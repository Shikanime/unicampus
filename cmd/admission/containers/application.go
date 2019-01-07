package containers

import (
	"context"

	"github.com/Shikanime/unicampus/cmd/admission/services/indexer"
	"github.com/Shikanime/unicampus/cmd/admission/services/persistence"
	"github.com/Shikanime/unicampus/pkg/admission"
	"github.com/Shikanime/unicampus/pkg/unicampus_admission"
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

func (s *Student) AppyStudentToSchool(ctx context.Context, in *unicampus_admission.Application) (*unicampus_admission.Application, error) {
	applicationData, err := s.persistence.CreateApplication(NewApplicationNetworkToDomain(in))
	if err != nil {
		return nil, err
	}

	return NewApplicationDomainToNetwork(applicationData), nil
}

func NewApplicationNetworkToDomain(application *unicampus_admission.Application) *admission.Application {
	return &admission.Application{
		UUID:    application.Uuid,
		School:  NewSchoolNetworkToDomain(application.School),
		Student: NewStudentNetworkToDomain(application.Student),
	}
}

func NewApplicationDomainToNetwork(application *admission.Application) *unicampus_admission.Application {
	return &unicampus_admission.Application{
		Uuid:    application.UUID,
		School:  NewSchoolDomainToNetwork(application.School),
		Student: NewStudentDomainToNetwork(application.Student),
	}
}
