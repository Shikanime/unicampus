package services

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type PostgreSQLService struct {
	driver *gorm.DB
}

func NewPostgreSQLService() *PostgreSQLService {
	conn, err := gorm.Open("postgres", "sslmode=disable user=postgres password=postgres dbname=yo")
	if err != nil {
		log.Fatalf("failed to connect postgres service: %v", err)
	}

	return &PostgreSQLService{
		driver: conn,
	}
}

func (s *PostgreSQLService) Driver() *gorm.DB {
	return s.driver
}

func (s *PostgreSQLService) Close() {
	if err := s.driver.Close(); err != nil {
		log.Fatalf("failed to close postgres service connection: %v", err)
	}
}
