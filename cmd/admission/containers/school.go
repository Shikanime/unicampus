package containers

import (
	"context"
	"io"

	"github.com/Shikanime/unicampus/cmd/admission/services/indexer"
	"github.com/Shikanime/unicampus/cmd/admission/services/persistence"
	"github.com/Shikanime/unicampus/pkg/admission"
	"github.com/Shikanime/unicampus/pkg/unicampus_admission"
)

func NewSchool(
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

func (s *School) ListSchools(stream unicampus_admission.AdmissionService_ListSchoolsServer) error {
	for {
		in, err := stream.Recv()

		if err == io.EOF {
			schoolData, err := s.persistence.GetSchool(NewSchoolNetworkToDomain(in))
			if err != nil {
				return err
			}

			if err := stream.Send(NewSchoolDomainToNetwork(schoolData)); err != nil {
				return err
			}
		}

		if err != nil {
			return err
		}

		return nil
	}
}

func (s *School) ListSchoolsByOffset(in *unicampus_admission.Offset, stream unicampus_admission.AdmissionService_ListSchoolsByOffsetServer) error {
	schoolDatas, err := s.persistence.ListSchoolsByOffset(in.First, in.Offset)

	if err != nil {
		return nil
	}

	for _, schoolData := range schoolDatas {
		if err := stream.Send(NewSchoolDomainToNetwork(schoolData)); err != nil {
			return err
		}
	}

	return nil
}

func (s *School) ListSchoolsByQuery(in *unicampus_admission.Query, stream unicampus_admission.AdmissionService_ListSchoolsByQueryServer) error {
	schoolIndexes, err := s.indexer.SearchSchoolsByQuery(in.Query)
	if err != nil {
		return err
	}

	schoolDatas, err := s.persistence.ListSchools(schoolIndexes)
	if err != nil {
		return err
	}

	for _, schoolData := range schoolDatas {
		if err := stream.Send(NewSchoolDomainToNetwork(schoolData)); err != nil {
			return err
		}
	}

	return nil
}

func (s *School) ListSchoolsByCritera(in *unicampus_admission.Critera, stream unicampus_admission.AdmissionService_ListSchoolsByCriteraServer) error {
	if err := stream.Send(&unicampus_admission.School{}); err != nil {
		return err
	}
	return nil
}

func (s *School) GetSchool(ctx context.Context, in *unicampus_admission.School) (*unicampus_admission.School, error) {
	schoolData, err := s.persistence.GetSchool(NewSchoolNetworkToDomain(in))
	if err != nil {
		return nil, err
	}

	return NewSchoolDomainToNetwork(schoolData), nil
}

func (s *School) PutSchool(ctx context.Context, in *unicampus_admission.School) (*unicampus_admission.School, error) {
	schoolData, err := s.persistence.PutSchool(NewSchoolNetworkToDomain(in))
	if err != nil {
		return nil, err
	}

	return NewSchoolDomainToNetwork(schoolData), nil
}

func (s *School) DeleteSchool(ctx context.Context, in *unicampus_admission.School) (*unicampus_admission.School, error) {
	schoolData, err := s.persistence.DeleteSchool(NewSchoolNetworkToDomain(in))
	if err != nil {
		return nil, err
	}

	return NewSchoolDomainToNetwork(schoolData), nil
}

func NewSchoolNetworkToDomain(school *unicampus_admission.School) *admission.School {
	return &admission.School{
		UUID:        school.Uuid,
		Name:        school.Name,
		Description: school.Description,
	}
}

func NewSchoolDomainToNetwork(school *admission.School) *unicampus_admission.School {
	return &unicampus_admission.School{
		Uuid:        school.UUID,
		Name:        school.Name,
		Description: school.Description,
	}
}
