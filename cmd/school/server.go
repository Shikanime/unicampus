package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"os"

	"github.com/Shikanime/unicampus/cmd/school/domain"
	"github.com/Shikanime/unicampus/cmd/school/indexer"
	"github.com/Shikanime/unicampus/cmd/school/persistence"
	"github.com/Shikanime/unicampus/cmd/school/recommandation"
	"github.com/Shikanime/unicampus/internal/app/school"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	persistence    *persistence.Repo
	indexer        *indexer.Repo
	recommandation *recommandation.Repo
}

func (s *Server) ListSchool(stream unicampus_school.SchoolService_ListSchoolServer) error {
	for {
		in, err := stream.Recv()

		if err == io.EOF {
			schoolData := s.persistence.GetSchool(newSchoolNetworkToDomain(in))

			if err := stream.Send(newSchoolDomainToNetwork(schoolData)); err != nil {
				return err
			}
		}

		if err != nil {
			return err
		}

		return nil
	}
}

func (s *Server) ListSchoolByOffset(in *unicampus_school.Offset, stream unicampus_school.SchoolService_ListSchoolByOffsetServer) error {
	schoolDatas := s.persistence.ListSchoolByOffset(in.First, in.Offset)

	for _, schoolData := range schoolDatas {
		if err := stream.Send(newSchoolDomainToNetwork(schoolData)); err != nil {
			return err
		}
	}

	return nil
}

func (s *Server) ListSchoolByQuery(in *unicampus_school.Query, stream unicampus_school.SchoolService_ListSchoolByQueryServer) error {
	schoolIndexes := s.indexer.SearchSchoolByQuery(in.Query)
	schoolDatas := s.persistence.ListSchool(schoolIndexes)

	for _, schoolData := range schoolDatas {
		if err := stream.Send(newSchoolDomainToNetwork(schoolData)); err != nil {
			return err
		}
	}

	return nil
}

func (s *Server) ListSchoolByCritera(in *unicampus_school.Critera, stream unicampus_school.SchoolService_ListSchoolByCriteraServer) error {
	if err := stream.Send(&unicampus_school.School{}); err != nil {
		return err
	}
	return nil
}

func (s *Server) GetSchool(ctx context.Context, in *unicampus_school.School) (*unicampus_school.School, error) {
	return newSchoolDomainToNetwork(s.persistence.GetSchool(newSchoolNetworkToDomain(in))), nil
}

func (s *Server) PutSchool(ctx context.Context, in *unicampus_school.School) (*unicampus_school.School, error) {
	return newSchoolDomainToNetwork(s.persistence.PutSchool(newSchoolNetworkToDomain(in))), nil
}

func (s *Server) DeleteSchool(ctx context.Context, in *unicampus_school.School) (*unicampus_school.School, error) {
	return newSchoolDomainToNetwork(s.persistence.DeleteSchool(newSchoolNetworkToDomain(in))), nil
}

func newSchoolNetworkToDomain(school *unicampus_school.School) *domain.School {
	return &domain.School{
		ID:          school.Id,
		Name:        school.Name,
		Description: school.Description,
	}
}

func newSchoolDomainToNetwork(school *domain.School) *unicampus_school.School {
	return &unicampus_school.School{
		Id:          school.ID,
		Name:        school.Name,
		Description: school.Description,
	}
}

func NewTCPListener() net.Listener {
	port := os.Getenv("PORT")
	if port == "" {
		port = "50051"
	}

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	return listener
}

func main() {
	tcpListener := NewTCPListener()

	persistenceRepo := persistence.NewClient()
	defer persistenceRepo.Close()
	indexerRepo := indexer.NewClient()
	defer indexerRepo.Close()
	recommandationRepo := recommandation.NewClient()
	defer recommandationRepo.Close()

	// Server
	s := grpc.NewServer()
	unicampus_school.RegisterSchoolServiceServer(s, &Server{
		persistence:    persistenceRepo,
		indexer:        indexerRepo,
		recommandation: recommandationRepo,
	})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(tcpListener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
