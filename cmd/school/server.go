package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/Shikanime/unicampus/cmd/school/indexer"
	"github.com/Shikanime/unicampus/cmd/school/persistence"
	"github.com/Shikanime/unicampus/cmd/school/recommandation"
	"github.com/Shikanime/unicampus/pkg/school"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	persistence    *persistence.Repo
	indexer        *indexer.Repo
	recommandation *recommandation.Repo
}

func (s *Server) SearchSchoolByQuery(ctx context.Context, in *unicampus_school.SearchSchoolByQueryRequest) (*unicampus_school.SearchSchoolByQueryReply, error) {
	schoolIndexes := s.indexer.SearchSchoolByQuery(in.Query)
	schoolDatas := s.persistence.ListSchool(schoolIndexes)

	schools := make([]*unicampus_school.School, len(schoolDatas))
	for i, dbSchool := range schoolDatas {
		schools[i] = &unicampus_school.School{
			Name:        dbSchool.Name,
			Description: dbSchool.Description,
		}
	}

	return &unicampus_school.SearchSchoolByQueryReply{
		Schools: schools,
	}, nil
}

func (s *Server) SearchSchoolByCritera(ctx context.Context, in *unicampus_school.SearchSchoolByCriteraRequest) (*unicampus_school.SearchSchoolByCriteraReply, error) {
	return &unicampus_school.SearchSchoolByCriteraReply{}, nil
}

func (s *Server) ListSchool(ctx context.Context, in *unicampus_school.ListSchoolRequest) (*unicampus_school.ListSchoolReply, error) {
	schoolDatas := s.persistence.ListSchoolByOffset(in.First, in.Offset)

	schools := make([]*unicampus_school.School, len(schoolDatas))
	for i, dbSchool := range schoolDatas {
		schools[i] = &unicampus_school.School{
			Name:        dbSchool.Name,
			Description: dbSchool.Description,
		}
	}

	return &unicampus_school.ListSchoolReply{
		Schools: schools,
	}, nil
}

func (s *Server) GetSchool(ctx context.Context, in *unicampus_school.GetSchoolRequest) (*unicampus_school.GetSchoolReply, error) {
	schoolData := s.persistence.GetSchool(in.Id)

	return &unicampus_school.GetSchoolReply{
		School: &unicampus_school.School{
			Name:        schoolData.Name,
			Description: schoolData.Description,
		},
	}, nil
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
