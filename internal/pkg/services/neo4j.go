package services

import (
	"log"

	neo4j "github.com/johnnadratowski/golang-neo4j-bolt-driver"
)

type Neo4jService struct {
	driver neo4j.Conn
}

func NewNeo4jService() *Neo4jService {
	driver := neo4j.NewDriver()
	conn, err := driver.OpenNeo("bolt://neo4j:neosecret@localhost:7687")
	if err != nil {
		log.Fatalf("failed to connect neo4j service")
	}

	return &Neo4jService{
		driver: conn,
	}
}

func (s *Neo4jService) Driver() neo4j.Conn {
	return s.driver
}

func (s *Neo4jService) Close() {
	if err := s.driver.Close(); err != nil {
		log.Fatalf("failed to close neo4j service connection: %v", err)
	}
}
