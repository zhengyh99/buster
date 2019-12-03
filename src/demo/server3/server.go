package main

import (
	"buster/iface"
	"buster/net"
	"buster/routers"
	"buster/utils"
	"fmt"
)

func OnConnBegin(conn iface.IConnection) {
	fmt.Println("/////// on conn begin is calling.....")
	conn.Send(202, []byte("Connection is established "))
}

func OnConnLost(conn iface.IConnection) {
	fmt.Println("///////// on conn lost is calling")
	fmt.Printf("conn id :%d is lost........", conn.GetConnID())
}
func main() {
	server := net.NewServer(utils.GlobalObject)
	server.SetOnConnStart(OnConnBegin)
	server.SetOnConnStop(OnConnLost)
	server.AddRouter(0, &routers.DataPackRouter{})
	server.AddRouter(1, &routers.PingRouter{})
	server.Run()
}
