package services

import (
	"fmt"
	"log"
	"os"
	"strconv"

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

func lookupElasticSearchPort() uint16 {
	env, ok := os.LookupEnv("ELASTICSEARCH_PORT")
	if !ok {
		return 9200
	}
	port, err := strconv.ParseUint(env, 10, 16)
	if err != nil {
		panic(err)
	}
	return uint16(port)
}

func NewElasticSearchService(name string) *ElasticSearchService {
	conn, err := elastic.NewClient(
		elastic.SetURL(fmt.Sprintf(
			"http://%s:%d",
			lookupElasticSearchHost(),
			lookupElasticSearchPort(),
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
