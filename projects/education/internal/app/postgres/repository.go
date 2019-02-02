package neo4j

import (
	neo4jDriver "github.com/johnnadratowski/golang-neo4j-bolt-driver"

	education_api_v1alpha1 "gitlab.com/deva-hub/unicampus/projects/education/api/v1alpha1"
	"gitlab.com/deva-hub/unicampus/projects/education/internal/pkg/services"
)

// TODO Separate to seed
const (
	regionDataModel = `
    CREATE (Region {city: "Paris", state: "France", country: "France",zipcode: "75000"})
    CREATE (Region {city: "Paris", state: "France", country: "France",zipcode: "75001"})
    CREATE (Region {city: "Paris", state: "France", country: "France",zipcode: "75002"})
    CREATE (Region {city: "Paris", state: "France", country: "France",zipcode: "75003"})
    CREATE (Region {city: "Paris", state: "France", country: "France",zipcode: "75004"})
    CREATE (Region {city: "Paris", state: "France", country: "France",zipcode: "75005"})
    CREATE (Region {city: "Paris", state: "France", country: "France",zipcode: "75006"})
    CREATE (Region {city: "Paris", state: "France", country: "France",zipcode: "75007"})
    CREATE (Region {city: "Paris", state: "France", country: "France",zipcode: "75008"})
    CREATE (Region {city: "Paris", state: "France", country: "France",zipcode: "75009"})
    CREATE (Region {city: "Paris", state: "France", country: "France",zipcode: "75010"})
    CREATE (Region {city: "Paris", state: "France", country: "France",zipcode: "75011"})
    CREATE (Region {city: "Paris", state: "France", country: "France",zipcode: "75012"})
    CREATE (Region {city: "Paris", state: "France", country: "France",zipcode: "75013"})
    CREATE (Region {city: "Paris", state: "France", country: "France",zipcode: "75014"})
    CREATE (Region {city: "Paris", state: "France", country: "France",zipcode: "75015"})
    CREATE (Region {city: "Paris", state: "France", country: "France",zipcode: "75016"})
    CREATE (Region {city: "Paris", state: "France", country: "France",zipcode: "75017"})
    CREATE (Region {city: "Paris", state: "France", country: "France",zipcode: "75018"})
    CREATE (Region {city: "Paris", state: "France", country: "France",zipcode: "75019"})
    CREATE (Region {city: "Paris", state: "France", country: "France",zipcode: "75020"})
    CREATE (Region {city: "Paris", state: "France", country: "France",zipcode: "75116"})
    CREATE (Region {city: "Paris", state: "France", country: "France",zipcode: "75680"})
    `
)

func New(service *services.Neo4jService) *Repository {
	return &Repository{
		conn: service.Driver(),
	}
}

type Repository struct {
	conn neo4jDriver.Conn
}

func (r *Repository) Init() error {
	// if _, err := r.conn.ExecNeo(regionDataModel, map[string]interface{}{}); err != nil {
	// 	return err
	// }
	return nil
}

func (r *Repository) RecommandSchoolsByCritera(school *education_api_v1alpha1.Critera) ([]*education_api_v1alpha1.School, error) {
	return nil, nil
}

func (r *Repository) PutSchool(school *education_api_v1alpha1.School) error {
	if _, err := r.conn.ExecNeo(`CREATE (School {id: id})`, map[string]interface{}{
		"id": school.UUID,
	}); err != nil {
		return err
	}
	return nil
}

func (r *Repository) DeleteSchool(school *education_api_v1alpha1.School) error {
	if _, err := r.conn.ExecNeo(`MATCH (s:School {id: id}) DELETE (s)`, map[string]interface{}{
		"id": school.UUID,
	}); err != nil {
		return err
	}
	return nil
}
