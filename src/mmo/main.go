package main

import (
	"buster/bnet"
	"buster/iface"
	"mmo/kernel"
)

func LoginInit(conn iface.IConnection) {
	player := kernel.NewPlayer(conn)
	player.SyncPid()
	player.BroadCastStartPosition()

}

func main() {
	server := bnet.NewServer()
	server.SetOnConnStart(LoginInit)
	server.Run()
}
