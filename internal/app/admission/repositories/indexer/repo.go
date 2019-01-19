package indexer

import (
	"context"
	"errors"

	"github.com/Shikanime/unicampus/internal/pkg/services"
	"github.com/Shikanime/unicampus/pkg/admission"
	"github.com/olivere/elastic"
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

	exists, err := r.conn.IndexExists(schoolIndexName).Do(ctx)
	if err != nil {
		return err
	}

	if !exists {
		// createIndex, err := r.conn.CreateIndex(schoolIndexName).
		// 	BodyString(mapping).
		// 	Do(ctx)
		// if err != nil {
		// 	return err
		// }

		// if !createIndex.Acknowledged {
		// 	return errors.New("Not acknowledged")
		// }
	}
	return nil
}

func (r *Repo) SearchSchools(school *admission.School) ([]*admission.School, error) {
	return nil, nil
}

func (r *Repo) SearchSchoolsByQuery(query string) ([]*admission.School, error) {
	queryBuilder := elastic.NewMultiMatchQuery(query, "name", "description").
		Fuzziness("2").
		MinimumShouldMatch("2")

	result, err := r.conn.Search().
		Index(schoolIndexName).
		Query(queryBuilder).
		Do(context.Background())
	if err != nil {
		return nil, err
	}

	return newSchoolsIndexerToDomain(result.Hits.Hits), nil
}

func (r *Repo) PutSchool(school *admission.School) error {
	result, err := r.conn.Index().
		Index(schoolIndexName).
		Type(schoolTypeName).
		Id(school.UUID).
		BodyJson(school).
		Do(context.Background())
	if err != nil {
		return err
	}
	if result.Result != "created" {
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
