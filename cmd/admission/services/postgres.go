package services

import (
	"log"

	"github.com/jinzhu/gorm"
)

type PostgresService struct {
	driver *gorm.DB
}

func NewPostgresService() *PostgresService {
	conn, err := gorm.Open("postgres", "sslmode=disable user=postgres password=postgres dbname=yo")
	if err != nil {
		log.Fatalf("failed to connect postgres service: %v", err)
	}

	// Migrate
	if err != nil {
		log.Fatalf("failed to connect postgres service: %v", err)
	}

	return &PostgresService{
		driver: conn,
	}
}

func (s *PostgresService) Driver() *gorm.DB {
	return s.driver
}

func (s *PostgresService) Close() {
	if err := s.driver.Close(); err != nil {
		log.Fatalf("failed to close postgres service connection: %v", err)
	}
}
