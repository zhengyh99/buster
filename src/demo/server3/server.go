package main

import (
	"buster/net"
	"buster/routers"
	"buster/utils"
)

func main() {
	server := net.NewServer(utils.GlobalObject)
	server.AddRouter(0,&routers.DataPackRouter{})
	server.AddRouter(1,&routers.PingRouter{})
	server.Run()
}
