package app

import (
	"context"
	"errors"

	unicampus_api_admission_v1alpha1 "github.com/Shikanime/unicampus/api/admission/v1alpha1"
	"github.com/Shikanime/unicampus/internal/app/admission/repositories/indexer"
	"github.com/Shikanime/unicampus/internal/app/admission/repositories/persistence"
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

func (s *Student) GetStudent(ctx context.Context, in *unicampus_api_admission_v1alpha1.Student) (*unicampus_api_admission_v1alpha1.Student, error) {
	return &unicampus_api_admission_v1alpha1.Student{}, errors.New("unmplemented")
}

func (s *Student) PutStudent(ctx context.Context, in *unicampus_api_admission_v1alpha1.Student) (*unicampus_api_admission_v1alpha1.Student, error) {
	return &unicampus_api_admission_v1alpha1.Student{}, errors.New("unmplemented")
}

func (s *Student) DeleteStudent(ctx context.Context, in *unicampus_api_admission_v1alpha1.Student) (*unicampus_api_admission_v1alpha1.Student, error) {
	return &unicampus_api_admission_v1alpha1.Student{}, errors.New("unmplemented")
}
