package services

import (
	"log"

	"github.com/johnnadratowski/golang-neo4j-bolt-driver"
)

type Neo4jService struct {
	driver golangNeo4jBoltDriver.Conn
}

func NewNeo4jService() *Neo4jService {
	driver := golangNeo4jBoltDriver.NewDriver()
	conn, err := driver.OpenNeo("bolt://neo4j:neosecret@localhost:7687")
	if err != nil {
		log.Fatalf("failed to connect neo4j service")
	}

	return &Neo4jService{
		driver: conn,
	}
}

func (s *Neo4jService) Driver() golangNeo4jBoltDriver.Conn {
	return s.driver
}

func (s *Neo4jService) Close() {
	if err := s.driver.Close(); err != nil {
		log.Fatalf("failed to close neo4j service connection: %v", err)
	}
}
