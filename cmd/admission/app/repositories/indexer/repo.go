package indexer

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/olivere/elastic"
	"gitlab.com/deva-hub/unicampus/internal/pkg/services"
	"gitlab.com/deva-hub/unicampus/pkg/admission"
)

func NewRepository(service *services.ElasticSearchService) Repo {
	return Repo{
		conn: service.Driver(),
	}
}

type Repo struct {
	conn *elastic.Client
}

func (r *Repo) Init() error {
	ctx := context.Background()

	exists, err := r.conn.IndexExists(schoolIndexName).
		Do(ctx)
	if err != nil {
		return err
	}

	if !exists {
		createIndex, err := r.conn.CreateIndex(schoolIndexName).
			BodyString(fmt.Sprintf(`
        {
          "mapping": {
            %s
          }
        }
        `,
				schoolMap,
			)).
			Do(ctx)
		if err != nil {
			return err
		}

		if !createIndex.Acknowledged {
			return errors.New("Not acknowledged")
		}
	}
	return nil
}

func (r *Repo) SearchSchool(school *admission.School) ([]*admission.School, error) {
	nameTermQuery := elastic.NewTermQuery("name", school.Name)
	descriptionTermQuery := elastic.NewTermQuery("description", school.Description)

	dbSchool, err := r.conn.Search().
		Index(schoolIndexName).
		Query(nameTermQuery).
		Query(descriptionTermQuery).
		Sort("user", true).
		Pretty(true).
		Do(context.Background())
	if err != nil {
		return nil, err
	}

	res := make([]*admission.School, len(dbSchool.Hits.Hits))
	for i, hit := range dbSchool.Hits.Hits {
		school := new(School)
		json.Unmarshal(*hit.Source, school)
		res[i] = formatSchoolDomain(*school)
	}

	return res, nil
}

func (r *Repo) SearchSchoolsByQuery(query string) ([]*admission.School, error) {
	queryBuilder := elastic.NewMultiMatchQuery(query, "name", "description").
		Fuzziness("2").
		MinimumShouldMatch("2")

	dbSchool, err := r.conn.Search().
		Index(schoolIndexName).
		Query(queryBuilder).
		Do(context.Background())
	if err != nil {
		return nil, err
	}

	res := make([]*admission.School, len(dbSchool.Hits.Hits))
	for i, hit := range dbSchool.Hits.Hits {
		school := new(School)
		json.Unmarshal(*hit.Source, school)
		res[i] = formatSchoolDomain(*school)
	}

	return res, nil
}

func (r *Repo) PutSchool(school *admission.School) error {
	dbSchool, err := r.conn.Index().
		Index(schoolIndexName).
		Type(schoolTypeName).
		Id(school.UUID).
		BodyJson(school).
		Do(context.Background())
	if err != nil {
		return err
	}
	if dbSchool.Result != "created" {
		return errors.New("fail to index school")
	}

	return nil
}

func (r *Repo) DeleteSchool(school *admission.School) error {
	deleteIndex, err := r.conn.DeleteIndex(schoolIndexName).Do(context.Background())
	if err != nil {
		return err
	}
	if !deleteIndex.Acknowledged {
		return errors.New("fail to index school")
	}
	return nil
}

func formatSchoolDomain(in School) *admission.School {
	return &admission.School{
		Identification: admission.Identification{
			UUID: in.UUID,
		},
	}
}
