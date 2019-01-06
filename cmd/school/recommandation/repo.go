package recommandation

import (
	bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
)

type Repo struct {
	Conn bolt.Conn
}
