package main

import (
	"context"
	"io"

	"gitlab.com/deva-hub/unicampus/internal/app/education/elasticsearch"
	"gitlab.com/deva-hub/unicampus/internal/app/education/neo4j"
	"gitlab.com/deva-hub/unicampus/internal/app/education/postgres"

	unicampus_api_education_v1alpha1 "gitlab.com/deva-hub/unicampus/api/v1alpha1"
	"gitlab.com/deva-hub/unicampus/internal/pkg/delivers"
	"gitlab.com/deva-hub/unicampus/internal/pkg/services"
)

func main() {
	grpcDeliver := delivers.NewGRPCDeliver()

	pg := services.NewPostgreSQLService("education")
	defer pg.Close()
	es := services.NewElasticSearchService("education")
	defer es.Close()
	nj := services.NewNeo4jService("education")
	defer es.Close()

	pgr := postgres.New(pg)
	esr := elasticsearch.New(es)
	njr := neo4j.New(nj)

	if err := pgr.Init(); err != nil {
		panic(err)
	}
	if err := esr.Init(); err != nil {
		panic(err)
	}
	if err := njr.Init(); err != nil {
		panic(err)
	}

	unicampus_api_education_v1alpha1.RegisterAdmissionServiceServer(grpcDeliver.Server(), &Server{
		storage:        pgr,
		search:         esr,
		recommandation: njr,
	})

	grpcDeliver.Run()
}

type Server struct {
	storage        *postgres.Repository
	search         *elasticsearch.Repository
	recommandation *neo4j.Repository
}

func (s *Server) ListSchools(stream unicampus_api_education_v1alpha1.AdmissionService_ListSchoolsServer) error {
	for {
		in, err := stream.Recv()

		if err == io.EOF {
			school, err := s.storage.GetSchool(in)
			if err != nil {
				return err
			}

			if err := stream.Send(school); err != nil {
				return err
			}
		}

		if err != nil {
			return err
		}

		return nil
	}
}

func (s *Server) ListSchoolsByQuery(in *unicampus_api_education_v1alpha1.Query, stream unicampus_api_education_v1alpha1.AdmissionService_ListSchoolsByQueryServer) error {
	schools, err := s.search.SearchSchoolsByQuery(in.Content)
	if err != nil {
		return err
	}

	schools, err = s.storage.ListSchools(schools)
	if err != nil {
		return err
	}

	for _, school := range schools {
		if err := stream.Send(school); err != nil {
			return err
		}
	}

	return nil
}

func (s *Server) ListSchoolsByCritera(in *unicampus_api_education_v1alpha1.Critera, stream unicampus_api_education_v1alpha1.AdmissionService_ListSchoolsByCriteraServer) error {
	schools, err := s.recommandation.RecommandSchoolsByCritera(in)
	if err != nil {
		return err
	}

	schools, err = s.storage.ListSchools(schools)
	if err != nil {
		return err
	}

	for _, school := range schools {
		if err := stream.Send(school); err != nil {
			return err
		}
	}

	return nil
}

func (s *Server) RegisterSchool(ctx context.Context, in *unicampus_api_education_v1alpha1.School) (*unicampus_api_education_v1alpha1.School, error) {
	if err := s.storage.CreateSchool(in); err != nil {
		return nil, err
	}

	school, err := s.storage.GetSchool(in)
	if err != nil {
		return nil, err
	}

	if err := s.search.PutSchool(school); err != nil {
		return nil, err
	}

	if err := s.recommandation.PutSchool(school); err != nil {
		return nil, err
	}

	return in, nil
}

func (s *Server) UpdateSchool(ctx context.Context, in *unicampus_api_education_v1alpha1.School) (*unicampus_api_education_v1alpha1.School, error) {
	if err := s.storage.UpdateSchool(in); err != nil {
		return nil, err
	}

	school, err := s.storage.GetSchool(in)
	if err != nil {
		return nil, err
	}

	if err := s.search.PutSchool(school); err != nil {
		return nil, err
	}

	if err := s.recommandation.PutSchool(school); err != nil {
		return nil, err
	}

	return in, nil
}

func (s *Server) UnregisterSchool(ctx context.Context, in *unicampus_api_education_v1alpha1.School) (*unicampus_api_education_v1alpha1.School, error) {
	if err := s.storage.DeleteSchool(in); err != nil {
		return nil, err
	}

	if err := s.search.DeleteSchool(in); err != nil {
		return nil, err
	}

	if err := s.recommandation.DeleteSchool(in); err != nil {
		return nil, err
	}

	return in, nil
}

func (s *Server) RegisterStudent(ctx context.Context, in *unicampus_api_education_v1alpha1.Student) (*unicampus_api_education_v1alpha1.Student, error) {
	return nil, nil
}

func (s *Server) UpdateStudent(ctx context.Context, in *unicampus_api_education_v1alpha1.Student) (*unicampus_api_education_v1alpha1.Student, error) {
	return nil, nil
}

func (s *Server) UnregisterStudent(ctx context.Context, in *unicampus_api_education_v1alpha1.Student) (*unicampus_api_education_v1alpha1.Student, error) {
	return nil, nil
}
