package services

import (
	"log"

	"github.com/olivere/elastic"
)

type ElasticSearchService struct {
	driver *elastic.Client
}

func NewElasticSearchService() *ElasticSearchService {
	conn, err := elastic.NewClient(
		elastic.SetURL("http://localhost:9200"),
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
