package main

import (
	"flag"
	"github.com/burhon94/bFilesSvc.git/cmd/app/server"
	"github.com/burhon94/bdi/pkg/di"
	"github.com/gorilla/mux"
	"net"
)

var (
	hostFlag = flag.String("host", "0.0.0.0", "Server host")
	portFlag = flag.String("port", "9999", "Server host")
)

func main() {
	flag.Parse()
	addr := net.JoinHostPort(*hostFlag, *portFlag)
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
