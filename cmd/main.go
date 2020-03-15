package main

import (
	"flag"
	"github.com/burhon94/bFilesSvc/cmd/app/server"
	"github.com/burhon94/bFilesSvc/pkg/servces"
	"github.com/burhon94/bdi/pkg/di"
	"github.com/gorilla/mux"
	"log"
	"net"
)

const (
	ENV_HOST = "HOST"
	ENV_PORT = "PORT"
)

var (
	hostFlag = flag.String("host", "", "Server host")
	portFlag = flag.String("port", "", "Server host")
)

func main() {
	flag.Parse()
	log.Print("Settings host")
	host, ok := servces.EnvOrFlag(ENV_HOST, hostFlag)
	if !ok {
		panic("can't setting host")
	}

	log.Print("Settings port")
	port, ok := servces.EnvOrFlag(ENV_PORT, portFlag)
	if !ok {
		panic("can't setting port")
	}
	log.Print("host and port setting successfully")

	addr := net.JoinHostPort(host, port)
	log.Printf("server will start on: %s", addr)
	start(addr)
}

func start(addr string) {
	container := di.NewContainer()

	if container.Provide(
		server.NewServer,
		mux.NewRouter,
	) != nil {
		panic("can't provide container")
	}

	container.Start()

	var appServer *server.Server
	container.Component(&appServer)
	appServer.Start(addr)
}
