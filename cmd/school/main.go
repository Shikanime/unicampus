package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/Shikanime/unicampus/cmd/school/indexer"
	"github.com/Shikanime/unicampus/cmd/school/persistence"
	"github.com/Shikanime/unicampus/cmd/school/recommandation"
	"github.com/Shikanime/unicampus/pkg/school"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
	"github.com/olivere/elastic"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

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

func NewPersistenceClient() *gorm.DB {
	conn, err := gorm.Open("sqlite3", "gorm.db")
	if err != nil {
		log.Fatalf("failed to connect persistent database: %v", err)
	}
	return conn
}

func NewIndexerClient() *elastic.Client {
	conn, err := elastic.NewClient(
		elastic.SetURL("http://localhost:9200"),
		elastic.SetSniff(false),
	)
	if err != nil {
		log.Fatalf("failed to connect indexer database: %v", err)
	}
	return conn
}

func NewRecommandationClient() bolt.Conn {
	driver := bolt.NewDriver()
	conn, err := driver.OpenNeo("bolt://localhost:7687")
	if err != nil {
		log.Fatalf("failed to connect recommandation database: %v", err)
	}
	return conn
}

func main() {
	tcpListener := NewTCPListener()

	persistenceConn := NewPersistenceClient()
	defer persistenceConn.Close()
	indexerConn := NewIndexerClient()
	recommandationConn := NewRecommandationClient()

	// Migrate
	persistenceConn.AutoMigrate(&persistence.School{})

	// Seed
	persistenceConn.Create(&persistence.School{ID: "1", Name: "ETNA", Description: "Desc"})
	// bulk := indexerConn.Bulk().
	// 	Index("schools").
	// 	Type("_doc")
	// bulk.Add(elastic.NewBulkIndexRequest().Id("1").Doc(&indexer.School{ID: "1", Name: "ETNA", Description: "Desc"}))

	// _, err = bulk.Do(context.Background())
	// if err != nil {
	// 	log.Fatalf("failed to connect indexer database: %v", err)
	// }

	// Server
	s := grpc.NewServer()
	school.RegisterSchoolServiceServer(s, &Server{
		Persistence: &persistence.Repo{
			Conn: persistenceConn,
		},
		Indexer: &indexer.Repo{
			Conn: indexerConn,
		},
		Recommandation: &recommandation.Repo{
			Conn: recommandationConn,
		},
	})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(tcpListener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
