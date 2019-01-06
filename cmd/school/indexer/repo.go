package indexer

import (
	"context"
	"encoding/json"

	"github.com/olivere/elastic"
)

type Repo struct {
	Conn *elastic.Client
}

func (r *Repo) SearchSchoolByQuery(query string) []string {
	queryBuilder := elastic.NewMultiMatchQuery(query, "name", "description").
		Fuzziness("2").
		MinimumShouldMatch("2")

	result, err := r.Conn.Search().
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
