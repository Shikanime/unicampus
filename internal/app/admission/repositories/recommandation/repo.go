package recommandation

import (
	neo4j "github.com/johnnadratowski/golang-neo4j-bolt-driver"
	"gitlab.com/deva-hub/unicampus/internal/pkg/services"
)

func NewRepository(service *services.Neo4jService) *Repo {
	return &Repo{
		conn: service.Driver(),
	}
}

type Repo struct {
	conn neo4j.Conn
}
