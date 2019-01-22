package admission

import (
	"context"
	"io"

	unicampus_api_admission_v1alpha1 "github.com/Shikanime/unicampus/api/admission/v1alpha1"
	"github.com/Shikanime/unicampus/internal/app/admission/repositories/indexer"
	"github.com/Shikanime/unicampus/internal/app/admission/repositories/persistence"
	"github.com/Shikanime/unicampus/pkg/objconv"
)

func NewSchoolService(
	persistence *persistence.Repo,
	indexer *indexer.Repo,
) School {
	return School{
		persistence: persistence,
		indexer:     indexer,
	}
}

type School struct {
	persistence *persistence.Repo
	indexer     *indexer.Repo
}

func (s *School) ListSchools(stream unicampus_api_admission_v1alpha1.AdmissionService_ListSchoolsServer) error {
	for {
		in, err := stream.Recv()

		if err == io.EOF {
			schoolData, err := s.persistence.GetSchool(objconv.FormatSchoolDomain(in))
			if err != nil {
				return err
			}

			if err := stream.Send(objconv.FormatSchoolNetwork(schoolData)); err != nil {
				return err
			}
		}

		if err != nil {
			return err
		}

		return nil
	}
}

func (s *School) ListSchoolsByQuery(in *unicampus_api_admission_v1alpha1.Query, stream unicampus_api_admission_v1alpha1.AdmissionService_ListSchoolsByQueryServer) error {
	schoolIndexes, err := s.indexer.SearchSchoolsByQuery(in.Query)
	if err != nil {
		return err
	}

	schoolDatas, err := s.persistence.ListSchools(schoolIndexes)
	if err != nil {
		return err
	}

	for _, schoolData := range schoolDatas {
		if err := stream.Send(objconv.FormatSchoolNetwork(schoolData)); err != nil {
			return err
		}
	}

	return nil
}

func (s *School) ListSchoolsByCritera(in *unicampus_api_admission_v1alpha1.Critera, stream unicampus_api_admission_v1alpha1.AdmissionService_ListSchoolsByCriteraServer) error {
	if err := stream.Send(&unicampus_api_admission_v1alpha1.School{}); err != nil {
		return err
	}
	return nil
}

func (s *School) FindSchool(ctx context.Context, in *unicampus_api_admission_v1alpha1.School) (*unicampus_api_admission_v1alpha1.School, error) {
	schoolData, err := s.persistence.GetSchool(objconv.FormatSchoolDomain(in))
	if err != nil {
		return nil, err
	}

	return objconv.FormatSchoolNetwork(schoolData), nil
}

func (s *School) RegisterSchool(ctx context.Context, in *unicampus_api_admission_v1alpha1.School) (*unicampus_api_admission_v1alpha1.School, error) {
	school := objconv.FormatSchoolDomain(in)
	if err := s.persistence.CreateSchool(school); err != nil {
		return nil, err
	}

	if err := s.indexer.PutSchool(school); err != nil {
		return nil, err
	}

	return in, nil
}

func (s *School) UnregisterSchool(ctx context.Context, in *unicampus_api_admission_v1alpha1.School) (*unicampus_api_admission_v1alpha1.School, error) {
	school := objconv.FormatSchoolDomain(in)
	if err := s.persistence.DeleteSchool(school); err != nil {
		return nil, err
	}

	if err := s.indexer.DeleteSchool(school); err != nil {
		return nil, err
	}

	return in, nil
}

func (s *School) UpdateSchool(ctx context.Context, in *unicampus_api_admission_v1alpha1.School) (*unicampus_api_admission_v1alpha1.School, error) {
	school := objconv.FormatSchoolDomain(in)
	if err := s.persistence.UpdateSchool(school); err != nil {
		return nil, err
	}

	if err := s.indexer.PutSchool(school); err != nil {
		return nil, err
	}

	return in, nil
}
