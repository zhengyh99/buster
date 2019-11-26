package main

import (
	"buster/net"
	"buster/routers"
	"buster/utils"
)

func main() {
	server := net.NewServer(utils.GlobalObject)
	server.AddRouter(&routers.DataPackRouter{})
	server.Run()
}
