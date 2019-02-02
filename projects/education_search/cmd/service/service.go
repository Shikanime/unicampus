package main

import (
	"context"

	"gitlab.com/deva-hub/unicampus/projects/education_search/internal/app/elasticsearch"

	education_searchapi_v1alpha1 "gitlab.com/deva-hub/unicampus/projects/education_search/api/v1alpha1"
	"gitlab.com/deva-hub/unicampus/projects/education_search/internal/pkg/delivers"
	"gitlab.com/deva-hub/unicampus/projects/education_search/internal/pkg/services"
)

func main() {
	grpcDeliver := delivers.NewGRPCDeliver()

	elasticsearchService := services.NewElasticSearchService("education")
	defer elasticsearchService.Close()

	elasticsearchRepository := elasticsearch.New(elasticsearchService)

	if err := elasticsearchRepository.Init(); err != nil {
		panic(err)
	}

	education_searchapi_v1alpha1.RegisterEducationSearchServiceServer(grpcDeliver.Server(), &Server{
		elasticsearchRepository: elasticsearchRepository,
	})

	grpcDeliver.Run()
}

type Server struct {
	elasticsearchRepository *elasticsearch.Repository
}

func (s *Server) ListSchoolsByQuery(in *education_searchapi_v1alpha1.Query, stream education_searchapi_v1alpha1.EducationSearchService_ListSchoolsByQueryServer) error {
	schools, err := s.elasticsearchRepository.SearchSchoolsByQuery(in.Content)
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

func (s *Server) RegisterSchool(ctx context.Context, in *education_searchapi_v1alpha1.School) (*education_searchapi_v1alpha1.School, error) {
	if err := s.elasticsearchRepository.PutSchool(in); err != nil {
		return nil, err
	}
	return in, nil
}

func (s *Server) UpdateSchool(ctx context.Context, in *education_searchapi_v1alpha1.School) (*education_searchapi_v1alpha1.School, error) {
	if err := s.elasticsearchRepository.PutSchool(in); err != nil {
		return nil, err
	}
	return in, nil
}

func (s *Server) UnregisterSchool(ctx context.Context, in *education_searchapi_v1alpha1.School) (*education_searchapi_v1alpha1.School, error) {
	if err := s.elasticsearchRepository.DeleteSchool(in); err != nil {
		return nil, err
	}
	return in, nil
}
