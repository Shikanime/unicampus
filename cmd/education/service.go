package main

import (
	"context"
	"io"

	"gitlab.com/deva-hub/unicampus/cmd/education/app/repositories/postgres"

	unicampus_api_education_v1alpha1 "gitlab.com/deva-hub/unicampus/api/education/v1alpha1"
	"gitlab.com/deva-hub/unicampus/internal/pkg/delivers"
	"gitlab.com/deva-hub/unicampus/internal/pkg/services"
)

func main() {
	grpcDeliver := delivers.NewGRPCDeliver()

	postgresService := services.NewPostgreSQLService("education")
	defer postgresService.Close()

	postgresRepo := postgres.NewPostgresRepository(postgresService)

	unicampus_api_education_v1alpha1.RegisterAdmissionServiceServer(grpcDeliver.Server(), &Server{
		storage: postgresRepo,
	})

	grpcDeliver.Run()
}

type Server struct {
	storage *postgres.PostgresRepository
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

}

func (s *Server) ListSchoolsByCritera(in *unicampus_api_education_v1alpha1.Critera, stream unicampus_api_education_v1alpha1.AdmissionService_ListSchoolsByCriteraServer) error {

}

func (s *Server) RegisterSchool(ctx context.Context, in *unicampus_api_education_v1alpha1.School) (*unicampus_api_education_v1alpha1.School, error) {

}

func (s *Server) UpdateSchool(ctx context.Context, in *unicampus_api_education_v1alpha1.School) (*unicampus_api_education_v1alpha1.School, error) {

}

func (s *Server) UnregisterSchool(ctx context.Context, in *unicampus_api_education_v1alpha1.SchoolFilter) (*unicampus_api_education_v1alpha1.School, error) {

}

func (s *Server) RegisterStudent(ctx context.Context, in *unicampus_api_education_v1alpha1.Student) (*unicampus_api_education_v1alpha1.Student, error) {

}

func (s *Server) UpdateStudent(ctx context.Context, in *unicampus_api_education_v1alpha1.Student) (*unicampus_api_education_v1alpha1.Student, error) {

}

func (s *Server) UnregisterStudent(ctx context.Context, in *unicampus_api_education_v1alpha1.Student) (*unicampus_api_education_v1alpha1.Student, error) {

}

func formatSchoolDTO(in *postgres.School) *unicampus_api_education_v1alpha1.School {
	return &unicampus_api_education_v1alpha1.School{
		UUID:        in.UUID,
		Name:        in.Name,
		Description: in.Description,
		Phone:       in.Phone,
		Email:       in.Email,
	}
}
