package admission

import (
	"context"

	unicampus_api_admission_v1alpha1 "gitlab.com/deva-hub/unicampus/api/admission/v1alpha1"
	"gitlab.com/deva-hub/unicampus/internal/app/admission/repositories/indexer"
	"gitlab.com/deva-hub/unicampus/internal/app/admission/repositories/persistence"
	"gitlab.com/deva-hub/unicampus/pkg/objconv"
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
	if err := s.persistence.CreateApplication(objconv.FormatApplicationDomain(in)); err != nil {
		return nil, err
	}

	return in, nil
}
