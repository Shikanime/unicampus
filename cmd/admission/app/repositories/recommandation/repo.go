package recommandation

import (
	neo4j "github.com/johnnadratowski/golang-neo4j-bolt-driver"
	"gitlab.com/deva-hub/unicampus/internal/pkg/services"
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

func NewRepository(service *services.Neo4jService) *Repo {
	return &Repo{
		conn: service.Driver(),
	}
}

type Repo struct {
	conn neo4j.Conn
}

func (r *Repo) Init() error {
	if _, err := r.conn.ExecNeo(regionDataModel, map[string]interface{}{}); err != nil {
		return err
	}
	return nil
}
