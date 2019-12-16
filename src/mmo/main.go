package main

import (
	"buster/bnet"
	"buster/iface"
	"fmt"
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

func LoginOut(conn iface.IConnection) {
	pid := conn.GetProperty("pid")
	player := kernel.WorldMng.GetPlayer(pid.(int32))
	if player == nil {
		fmt.Println("worldmng getplayer error")
		return
	}
	player.OffLine()

}

func main() {
	server := bnet.NewServer()
	server.SetOnConnStart(LoginInit)
	server.SetOnConnStop(LoginOut)
	server.AddRouter(2, &apis.WorldChart{})
	server.AddRouter(3, &apis.MoveAction{})
	server.Run()
}
