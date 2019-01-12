package services

import (
	"log"

	"github.com/jinzhu/gorm"
)

type PostgreSQLDatabaseService struct {
	driver *gorm.DB
}

func NewPostgreSQLDatabaseService() *PostgreSQLDatabaseService {
	conn, err := gorm.Open("postgres", "sslmode=disable user=postgres password=postgres dbname=yo")
	if err != nil {
		log.Fatalf("failed to connect postgres service: %v", err)
	}

	// Migrate
	if err != nil {
		log.Fatalf("failed to connect postgres service: %v", err)
	}

	return &PostgreSQLDatabaseService{
		driver: conn,
	}
}

func (s *PostgreSQLDatabaseService) Get(out interface{}, clauses ...interface{}) error {
	if err := s.driver.Take(out, clauses).Error; err != nil {
		return err
	}
	return nil
}

func (s *PostgreSQLDatabaseService) Create(value interface{}) error {
	if err := s.driver.Create(value).Error; err != nil {
		return err
	}
	return nil
}

func (s *PostgreSQLDatabaseService) Update(value interface{}) error {
	if err := s.driver.Update(value).Error; err != nil {
		return err
	}
	return nil
}

func (s *PostgreSQLDatabaseService) Delete(value interface{}) error {
	if err := s.driver.Delete(value).Error; err != nil {
		return err
	}
	return nil
}

func (s *PostgreSQLDatabaseService) Driver() *gorm.DB {
	return s.driver
}

func (s *PostgreSQLDatabaseService) Close() {
	if err := s.driver.Close(); err != nil {
		log.Fatalf("failed to close postgres service connection: %v", err)
	}
}
