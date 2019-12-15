package kernel

import (
	"buster/iface"
	"errors"
	"fmt"
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
func (p *Player) Talk(content string) {
	pMsg := &pb.BroadCast{
		Pid: p.Pid,
		Tp:  1,
		Data: &pb.BroadCast_Content{
			Content: content,
		},
	}
	players := WorldMng.GetAllPlayer()

	for _, player := range players {
		player.SendMsg(200, pMsg)
	}
}
func (p *Player) getSurroundingPlayers() []*Player {
	//根据当前玩家的位置信息，找到周边玩家的集合pids
	pids := WorldMng.AoiMng.GetPlayersByPosition(p.X, p.Z)
	fmt.Println("pids:", pids)
	//周边玩家的集合players
	players := make([]*Player, 0, len(pids))

	for _, pid := range pids {
		players = append(players, WorldMng.GetPlayer(pid))
	}
	return players
}

//同步周围玩家
func (p *Player) SyncSurrounding() {
	players := p.getSurroundingPlayers()

	//组织广播信息

	fmt.Printf("player id: %d,x=%v,y=%v,z=%v,v=%v\n", p.Pid, p.X, p.Y, p.Z, p.V)

	myPMsg := &pb.BroadCast{
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

	playersPMsg := make([]*pb.Player, 0, len(players))

	for _, player := range players {
		player.SendMsg(200, myPMsg)
		pMsg := &pb.Player{
			Pid: player.Pid,
			P: &pb.Position{
				X: player.X,
				Y: player.Y,
				Z: player.Z,
				V: player.V,
			},
		}
		playersPMsg = append(playersPMsg, pMsg)
	}
	SycnPMsg := &pb.SyncPlayer{
		Ps: playersPMsg[:],
	}
	p.SendMsg(202, SycnPMsg)

}

func (p *Player) UpdatePosition(x, y, z, v float32) {
	p.X = x
	p.Y = y
	p.Z = z
	p.V = v

	pMsg := &pb.BroadCast{
		Pid: p.Pid,
		Tp:  4,
		Data: &pb.BroadCast_P{
			P: &pb.Position{
				X: p.X,
				Y: p.Y,
				Z: p.Z,
				V: p.V,
			},
		},
	}
	players := p.getSurroundingPlayers()
	for _, player := range players {
		player.SendMsg(200, pMsg)
	}
}
