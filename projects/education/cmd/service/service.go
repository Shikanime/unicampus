package main

import (
	"context"

	"gitlab.com/deva-hub/unicampus/projects/education/internal/app/postgres"
	education_api_v1alpha1 "gitlab.com/deva-hub/unicampus/projects/education/api/v1alpha1"
	"gitlab.com/deva-hub/unicampus/projects/education/internal/pkg/delivers"
	"gitlab.com/deva-hub/unicampus/projects/education/internal/pkg/services"
)

func main() {
	grpcDeliver := delivers.NewGRPCDeliver()

	postgresService := services.NewPostgreSQLService("education")
	defer pg.Close()


	postgresRepository := postgres.New(postgresService)

	if err := postgresRepository.Init(); err != nil {
		panic(err)
	}

	education_api_v1alpha1.RegisterEducationRecommandationServiceServer(grpcDeliver.Server(), &Server{
		postgresRepository: postgresRepository,
	})

	grpcDeliver.Run()
}

type Server struct {
	postgresRepository *postgres.Repository
}

func (s *Server) ListSchoolsByCritera(in *education_api_v1alpha1.Critera, stream education_api_v1alpha1.EducationRecommandationService_ListSchoolsByCriteraServer) error {
	schools, err := s.postgresRepository.RecommandSchoolsByCritera(in)
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

func (s *Server) RegisterSchool(ctx context.Context, in *education_api_v1alpha1.School) (*education_api_v1alpha1.School, error) {
	if err := s.postgresRepository.PutSchool(in); err != nil {
		return nil, err
	}

	return in, nil
}

func (s *Server) UpdateSchool(ctx context.Context, in *education_api_v1alpha1.School) (*education_api_v1alpha1.School, error) {
	if err := s.postgresRepository.PutSchool(in); err != nil {
		return nil, err
	}

	return in, nil
}

func (s *Server) UnregisterSchool(ctx context.Context, in *education_api_v1alpha1.School) (*education_api_v1alpha1.School, error) {
	if err := s.postgresRepository.DeleteSchool(in); err != nil {
		return nil, err
	}

	return in, nil
}
