package elasticsearch

import (
	"context"
	"encoding/json"
	"errors"

	elasticsearchDriver "github.com/olivere/elastic"
	unicampus_api_education_v1alpha1 "gitlab.com/deva-hub/unicampus/api/v1alpha1"
	"gitlab.com/deva-hub/unicampus/internal/pkg/services"
)

func New(service *services.ElasticSearchService) *Repository {
	return &Repository{
		conn: service.Driver(),
	}
}

type Repository struct {
	conn *elasticsearchDriver.Client
}

func (r *Repository) Init() error {
	ctx := context.Background()

	exists, err := r.conn.IndexExists(schoolIndexName).
		Do(ctx)
	if err != nil {
		return err
	}

	if !exists {
		// TODO: Fix unknow mapping property
		// createIndex, err := r.conn.CreateIndex(schoolIndexName).
		// 	BodyString(fmt.Sprintf(`
		//     {
		//       "mapping": {
		//         %s
		//       }
		//     }
		//     `,
		// 		schoolMap,
		// 	)).
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

func (r *Repository) SearchSchool(school *unicampus_api_education_v1alpha1.School) ([]*unicampus_api_education_v1alpha1.School, error) {
	nameTermQuery := elasticsearchDriver.NewTermQuery("name", school.Name)
	descriptionTermQuery := elasticsearchDriver.NewTermQuery("description", school.Description)

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

	res := make([]*unicampus_api_education_v1alpha1.School, len(dbSchool.Hits.Hits))
	for i, hit := range dbSchool.Hits.Hits {
		school := new(School)
		json.Unmarshal(*hit.Source, school)
		res[i] = formatSchoolDomain(school)
	}

	return res, nil
}

func (r *Repository) SearchSchoolsByQuery(query string) ([]*unicampus_api_education_v1alpha1.School, error) {
	queryBuilder := elasticsearchDriver.NewMultiMatchQuery(query, "name", "description").
		Fuzziness("2").
		MinimumShouldMatch("2")

	dbSchool, err := r.conn.Search().
		Index(schoolIndexName).
		Query(queryBuilder).
		Do(context.Background())
	if err != nil {
		return nil, err
	}

	res := make([]*unicampus_api_education_v1alpha1.School, len(dbSchool.Hits.Hits))
	for i, hit := range dbSchool.Hits.Hits {
		school := new(School)
		json.Unmarshal(*hit.Source, school)
		res[i] = formatSchoolDomain(school)
	}

	return res, nil
}

func (r *Repository) PutSchool(school *unicampus_api_education_v1alpha1.School) error {
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

func (r *Repository) DeleteSchool(school *unicampus_api_education_v1alpha1.School) error {
	deleteIndex, err := r.conn.DeleteIndex(schoolIndexName).Do(context.Background())
	if err != nil {
		return err
	}
	if !deleteIndex.Acknowledged {
		return errors.New("fail to index school")
	}
	return nil
}

func formatSchoolDomain(in *School) *unicampus_api_education_v1alpha1.School {
	return &unicampus_api_education_v1alpha1.School{
		UUID: in.UUID,
	}
}
