package indexer

import (
	"encoding/json"

	"github.com/Shikanime/unicampus/cmd/school/domain"
	"github.com/olivere/elastic"
)

type School struct {
	ID          string
	Name        string `json:"name"`
	Description string `json:"description"`
}

func newSchoolsIndexerToDomain(d []*elastic.SearchHit) []*domain.School {
	schoolIndexes := make([]*domain.School, len(d))
	for i, hit := range d {
		var school *School
		json.Unmarshal(*hit.Source, school)
		schoolIndexes[i] = &domain.School{
			ID: school.ID,
		}
	}
	return schoolIndexes
}
