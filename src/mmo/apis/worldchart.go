package apis

import (
	"buster/bnet"
	"buster/iface"
	"mmo/kernel"
	"mmo/pb"

	"github.com/golang/protobuf/proto"
)

type WorldChart struct {
	bnet.BaseRouter
}

func (wc *WorldChart) Handle(r iface.IRequest) {

	pmsg := &pb.Talk{}
	proto.Unmarshal(r.GetData(), pmsg)
	pid := r.GetConnection().GetProperty("pid")
	player := kernel.WorldMng.GetPlayer(pid.(int32))

	player.Talk(pmsg.Content)
}
