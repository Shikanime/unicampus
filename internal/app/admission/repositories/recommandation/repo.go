package recommandation

import (
	"github.com/Shikanime/unicampus/internal/pkg/services"
	neo4j "github.com/johnnadratowski/golang-neo4j-bolt-driver"
)

func NewRepository(service *services.Neo4jService) *Repo {
	return &Repo{
		conn: service.Driver(),
	}
}

type Repo struct {
	conn neo4j.Conn
}
