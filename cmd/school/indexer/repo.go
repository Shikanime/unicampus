package indexer

import (
	"context"
	"encoding/json"
	"log"

	"github.com/olivere/elastic"
)

func NewClient() *Repo {
	conn, err := elastic.NewClient(
		elastic.SetURL("http://localhost:9200"),
		elastic.SetSniff(false),
	)
	if err != nil {
		log.Fatalf("failed to connect indexer database: %v", err)
	}

	// bulk := conn.Bulk().
	// 	Index("schools").
	// 	Type("_doc")
	// bulk.Add(elastic.NewBulkIndexRequest().Id("1").Doc(&School{ID: "1", Name: "ETNA", Description: "Desc"}))

	// _, err = bulk.Do(context.Background())
	// if err != nil {
	// 	log.Fatalf("failed to connect indexer database: %v", err)
	// }

	return &Repo{
		conn: conn,
	}
}

type Repo struct {
	conn *elastic.Client
}

func (r *Repo) SearchSchoolByQuery(query string) []string {
	queryBuilder := elastic.NewMultiMatchQuery(query, "name", "description").
		Fuzziness("2").
		MinimumShouldMatch("2")

	result, err := r.conn.Search().
		Index("schools").
		Query(queryBuilder).
		Do(context.Background())
	if err != nil {
		panic(err)
	}

	schoolIndexes := make([]string, len(result.Hits.Hits))
	for i, hit := range result.Hits.Hits {
		var school School
		json.Unmarshal(*hit.Source, &school)
		schoolIndexes[i] = school.ID
	}

	return schoolIndexes
}

func (r *Repo) Close() {
}
