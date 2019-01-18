package services

import (
	"fmt"
	"log"
	"os"

	neo4j "github.com/johnnadratowski/golang-neo4j-bolt-driver"
)

type Neo4jService struct {
	driver neo4j.Conn
}

func lookupNeo4jUsername() string {
	username, ok := os.LookupEnv("NEO4J_USERNAME")
	if !ok {
		username = "postgres"
	}
	return username
}

func lookupNeo4jPassword() string {
	password, ok := os.LookupEnv("NEO4J_PASSWORD")
	if !ok {
		password = "unicampus"
	}
	return password
}

func lookupNeo4jHost() string {
	host, ok := os.LookupEnv("NEO4J_HOST")
	if !ok {
		host = "localhost"
	}
	return host
}

func NewNeo4jService() *Neo4jService {
	driver := neo4j.NewDriver()
	conn, err := driver.OpenNeo(fmt.Sprintf(
		"bolt://%s:%s@%s:7687",
		lookupNeo4jUsername(),
		lookupNeo4jPassword(),
		lookupNeo4jHost(),
	))
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
