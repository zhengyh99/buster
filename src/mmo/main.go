package main

import (
	"buster/bnet"
	"buster/iface"
	"mmo/apis"
	"mmo/kernel"
)

func LoginInit(conn iface.IConnection) {
	player := kernel.NewPlayer(conn)
	player.SyncPid()
	player.BroadCastStartPosition()
	kernel.WorldMng.AddPlayer(player)

	conn.SetProperty("pid", player.Pid)
	player.SyncSurrounding()

}

func main() {
	server := bnet.NewServer()
	server.SetOnConnStart(LoginInit)
	server.AddRouter(2, &apis.WorldChart{})
	server.AddRouter(3, &apis.MoveAction{})
	server.Run()
}
