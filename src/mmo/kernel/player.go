package kernel

import (
	"buster/iface"
	"errors"
	"math/rand"
	"mmo/pb"
	"sync"

	"github.com/golang/protobuf/proto"
)

//玩家对象
type Player struct {
	Pid  int32             //玩家id
	Conn iface.IConnection //与客户端链接
	X    float32           //玩家x轴
	Y    float32           //玩家垂直高度
	Z    float32           //玩家y轴
	V    float32           //玩家旋转角度
}

//PID 生成器
var PidGen int32 = 1

//PID 生成器锁
var PGLock sync.RWMutex

func NewPlayer(conn iface.IConnection) *Player {
	PGLock.Lock()
	pid := PidGen
	PidGen++
	PGLock.Unlock()
	return &Player{
		Pid:  pid,
		Conn: conn,
		X:    float32(160 + rand.Intn(10)),
		Y:    0,
		Z:    float32(140 + rand.Intn(20)),
		V:    0,
	}
}

func (p *Player) SendMsg(msgID uint32, data proto.Message) error {
	pdata, err := proto.Marshal(data)
	if err != nil {
		return err
	}
	if p.Conn == nil {
		return errors.New("Disconnected from client")
	}
	return p.Conn.Send(msgID, pdata)
}

//告知玩家 Pid ,同步服务器端的 pid 给客户端
func (p *Player) SyncPid() {
	pMsg := &pb.SyncPid{
		Pid: p.Pid,
	}
	//将消息发送给客户端
	p.SendMsg(1, pMsg)
}
func (p *Player) BroadCastStartPosition() {
	pMsg := &pb.BroadCast{
		Pid: p.Pid,
		Tp:  2,
		Data: &pb.BroadCast_P{
			P: &pb.Position{
				X: p.X,
				Y: p.Y,
				Z: p.Z,
				V: p.V,
			},
		},
	}
	p.SendMsg(200, pMsg)
}
