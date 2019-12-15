package apis

import (
	"buster/bnet"
	"buster/iface"
	"fmt"
	"mmo/kernel"
	"mmo/pb"

	"github.com/golang/protobuf/proto"
)

type MoveAction struct {
	bnet.BaseRouter
}

func (m *MoveAction) Handle(r iface.IRequest) {
	pMsg := &pb.Position{}

	err := proto.Unmarshal(r.GetData(), pMsg)
	if err != nil {
		fmt.Println("proto un marshal error:", err)
		return
	}

	pid := r.GetConnection().GetProperty("pid")

	fmt.Printf("Player id:=%d,Moved(%f,%f,%f,%f)", pid, pMsg.X, pMsg.Y, pMsg.Z, pMsg.V)

	player := kernel.WorldMng.GetPlayer(pid.(int32))
	if player == nil {
		fmt.Printf("Player id= %d is not existed", pid.(int32))
	}
	player.UpdatePosition(pMsg.X, pMsg.Y, pMsg.Z, pMsg.V)
}
