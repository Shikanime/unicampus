package recommandation

import (
	"log"

	bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
)

func NewClient() *Repo {
	driver := bolt.NewDriver()
	conn, err := driver.OpenNeo("bolt://localhost:7687")
	if err != nil {
		log.Fatalf("failed to connect recommandation database: %v", err)
	}

	return &Repo{
		conn: conn,
	}
}

type Repo struct {
	conn bolt.Conn
}

func (r *Repo) Close() {
}
