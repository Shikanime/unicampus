package admission

import (
	"context"
	"io"

	unicampus_api_admission_v1alpha1 "gitlab.com/deva-hub/unicampus/api/admission/v1alpha1"
	"gitlab.com/deva-hub/unicampus/internal/app/admission/repositories/indexer"
	"gitlab.com/deva-hub/unicampus/internal/app/admission/repositories/persistence"
	"gitlab.com/deva-hub/unicampus/pkg/objconv"
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
			school, err := s.persistence.GetSchool(objconv.FormatSchoolDomain(in))
			if err != nil {
				return err
			}

			if err := stream.Send(objconv.FormatSchoolNetwork(school)); err != nil {
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
	schools, err := s.indexer.SearchSchoolsByQuery(in.Query)
	if err != nil {
		return err
	}

	schools, err = s.persistence.ListSchools(schools)
	if err != nil {
		return err
	}

	for _, school := range schools {
		if err := stream.Send(objconv.FormatSchoolNetwork(school)); err != nil {
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
	school, err := s.persistence.GetSchool(objconv.FormatSchoolDomain(in))
	if err != nil {
		return nil, err
	}

	return objconv.FormatSchoolNetwork(school), nil
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
