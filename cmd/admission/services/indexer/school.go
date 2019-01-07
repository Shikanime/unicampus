package indexer

import (
	"encoding/json"

	"github.com/Shikanime/unicampus/pkg/admission"
	"github.com/olivere/elastic"
)

type School struct {
	UUID        string
	Name        string `json:"name"`
	Description string `json:"description"`
}

func newSchoolsIndexerToDomain(d []*elastic.SearchHit) []*admission.School {
	schoolIndexes := make([]*admission.School, len(d))
	for i, hit := range d {
		var school *School
		json.Unmarshal(*hit.Source, school)
		schoolIndexes[i] = &admission.School{
			UUID: school.UUID,
		}
	}
	return schoolIndexes
}
