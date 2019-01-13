package services

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type PostgreSQLService struct {
	driver *gorm.DB
}

func lookupPostgreSQLUsername(name string) string {
	username, ok := os.LookupEnv("POSTGRESQL_USERNAME")
	if !ok {
		username = fmt.Sprint("unicampus_", name)
	}
	return username
}

func lookupPostgreSQLPassword() string {
	password, ok := os.LookupEnv("POSTGRESQL_PASSWORD")
	if !ok {
		password = fmt.Sprint("unicampus_postgres")
	}
	return password
}

func lookupPostgreSQLHost() string {
	host, ok := os.LookupEnv("POSTGRESQL_HOST")
	if !ok {
		host = "localhost"
	}
	return host
}

func NewPostgreSQLService(name string) *PostgreSQLService {
	conn, err := gorm.Open("postgres", fmt.Sprintf(
		"sslmode=disable dbname=%s hostname=%s user=%s password=%s",
		name,
		lookupPostgreSQLHost(),
		lookupPostgreSQLUsername(name),
		lookupPostgreSQLPassword(),
	))
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
