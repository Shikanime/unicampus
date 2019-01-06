package main

import (
	"context"

	"github.com/Shikanime/unicampus/cmd/school/indexer"
	"github.com/Shikanime/unicampus/cmd/school/persistence"
	"github.com/Shikanime/unicampus/cmd/school/recommandation"
	"github.com/Shikanime/unicampus/pkg/school"
)

type Server struct {
	Persistence    *persistence.Repo
	Indexer        *indexer.Repo
	Recommandation *recommandation.Repo
}

func (s *Server) SearchSchoolByQuery(ctx context.Context, in *school.SearchSchoolByQueryRequest) (*school.SearchSchoolByQueryReply, error) {
	schoolIndexes := s.Indexer.SearchSchoolByQuery(in.Query)
	schoolDatas := s.Persistence.ListSchool(schoolIndexes)

	schools := make([]*school.School, len(schoolDatas))
	for i, dbSchool := range schoolDatas {
		schools[i] = &school.School{
			Name:        dbSchool.Name,
			Description: dbSchool.Description,
		}
	}

	return &school.SearchSchoolByQueryReply{
		Schools: schools,
	}, nil
}

func (s *Server) SearchSchoolByCritera(ctx context.Context, in *school.SearchSchoolByCriteraRequest) (*school.SearchSchoolByCriteraReply, error) {
	return &school.SearchSchoolByCriteraReply{}, nil
}

func (s *Server) ListSchool(ctx context.Context, in *school.ListSchoolRequest) (*school.ListSchoolReply, error) {
	schoolDatas := s.Persistence.ListSchoolByOffset(in.First, in.Offset)

	schools := make([]*school.School, len(schoolDatas))
	for i, dbSchool := range schoolDatas {
		schools[i] = &school.School{
			Name:        dbSchool.Name,
			Description: dbSchool.Description,
		}
	}

	return &school.ListSchoolReply{
		Schools: schools,
	}, nil
}

func (s *Server) GetSchool(ctx context.Context, in *school.GetSchoolRequest) (*school.GetSchoolReply, error) {
	schoolData := s.Persistence.GetSchool(in.Id)

	return &school.GetSchoolReply{
		School: &school.School{
			Name:        schoolData.Name,
			Description: schoolData.Description,
		},
	}, nil
}
