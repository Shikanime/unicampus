package main

import (
	"context"

	"gitlab.com/deva-hub/unicampus/projects/education_recommandation/internal/app/neo4j"

	education_api_v1alpha1 "gitlab.com/deva-hub/unicampus/projects/education_recommandation/api/v1alpha1"
	"gitlab.com/deva-hub/unicampus/projects/education_recommandation/internal/pkg/delivers"
	"gitlab.com/deva-hub/unicampus/projects/education_recommandation/internal/pkg/services"
)

func main() {
	grpcDeliver := delivers.NewGRPCDeliver()

	neo4jService := services.NewNeo4jService("education")
	defer neo4jService.Close()

	neo4jRepository := neo4j.New(neo4jService)

	if err := neo4jRepository.Init(); err != nil {
		panic(err)
	}

	education_api_v1alpha1.RegisterEducationRecommandationServiceServer(grpcDeliver.Server(), &Server{
		neo4jRepository: neo4jRepository,
	})

	grpcDeliver.Run()
}

type Server struct {
	neo4jRepository *neo4j.Repository
}

func (s *Server) ListSchoolsByCritera(in *education_api_v1alpha1.Critera, stream education_api_v1alpha1.EducationRecommandationService_ListSchoolsByCriteraServer) error {
	schools, err := s.neo4jRepository.RecommandSchoolsByCritera(in)
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
	if err := s.neo4jRepository.PutSchool(in); err != nil {
		return nil, err
	}

	return in, nil
}

func (s *Server) UpdateSchool(ctx context.Context, in *education_api_v1alpha1.School) (*education_api_v1alpha1.School, error) {
	if err := s.neo4jRepository.PutSchool(in); err != nil {
		return nil, err
	}

	return in, nil
}

func (s *Server) UnregisterSchool(ctx context.Context, in *education_api_v1alpha1.School) (*education_api_v1alpha1.School, error) {
	if err := s.neo4jRepository.DeleteSchool(in); err != nil {
		return nil, err
	}

	return in, nil
}
