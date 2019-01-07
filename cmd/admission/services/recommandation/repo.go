package recommandation

import (
	"log"

	"github.com/johnnadratowski/golang-neo4j-bolt-driver"
)

func NewClient() *Repo {
	driver := golangNeo4jBoltDriver.NewDriver()
	conn, err := driver.OpenNeo("bolt://neo4j:neosecret@localhost:7687")
	if err != nil {
		log.Fatalf("failed to connect recommandation database")
	}

	return &Repo{
		conn: conn,
	}
}

type Repo struct {
	conn golangNeo4jBoltDriver.Conn
}

func (r *Repo) Close() {
	if err := r.conn.Close(); err != nil {
		panic(err)
	}
}
