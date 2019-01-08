package delivers

import (
	"log"

	"github.com/kataras/iris"
)

type HTTPDeliver struct {
	server *iris.Application
}

func NewHTTPDeliver() *HTTPDeliver {
	server := iris.Default()

	return &HTTPDeliver{
		server: server,
	}
}

func (d *HTTPDeliver) Server() *iris.Application {
	return d.server
}

func (d *HTTPDeliver) Run() {
	if err := d.server.Run(iris.Addr(":8080")); err != nil {
		log.Fatalf("")
	}
}
