package services

import (
	"fmt"
	"log"
	"os"

	"github.com/olivere/elastic"
)

type ElasticSearchService struct {
	driver *elastic.Client
}

func lookupElasticSearchHost() string {
	host, ok := os.LookupEnv("ELASTICSEARCH_HOST")
	if !ok {
		host = "localhost"
	}
	return host
}

func NewElasticSearchService(name string) *ElasticSearchService {
	conn, err := elastic.NewClient(
		elastic.SetURL(fmt.Sprintf(
			"http://%s:9200",
			lookupElasticSearchHost(),
		)),
		elastic.SetSniff(false),
	)
	if err != nil {
		log.Fatalf("failed to connect indexer database: %v", err)
	}

	return &ElasticSearchService{
		driver: conn,
	}
}

func (s *ElasticSearchService) Driver() *elastic.Client {
	return s.driver
}

func (r *ElasticSearchService) Close() {
}
