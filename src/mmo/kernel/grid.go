package kernel

import (
	"fmt"
	"sync"
)

type Grid struct {
	//id
	GID int
	//左坐标
	Left int
	//右坐标
	Right int
	//上坐标
	Top int
	//下坐标
	Bottom int
	//玩家集合
	Players map[int]bool
	//玩家集合读写锁
	PlayersLock sync.RWMutex
}

//新建格式
func NewGrid(gid, left, right, top, bottom int) *Grid {
	return &Grid{
		GID:     gid,
		Left:    left,
		Right:   right,
		Top:     top,
		Bottom:  bottom,
		Players: make(map[int]bool),
	}
}

//添加玩家
func (g *Grid) Add(playerID int) {
	g.PlayersLock.Lock()
	defer g.PlayersLock.Unlock()
	g.Players[playerID] = true
}

//删除玩家
func (g *Grid) Remove(playerID int) {
	g.PlayersLock.Lock()
	defer g.PlayersLock.Unlock()
	delete(g.Players, playerID)
}

//返回所有玩家
func (g *Grid) GetAll() (players []int) {
	g.PlayersLock.RLock()
	defer g.PlayersLock.RUnlock()
	for key, _ := range g.Players {
		players = append(players, key)
	}
	return
}

//返回格式信息
func (g *Grid) String() string {
	return fmt.Sprintf("Gid:=%d,Left:%d,Right:%d,Top:%d,Bottom:%d,Players:%v",
		g.GID, g.Left, g.Right, g.Top, g.Bottom, g.Players)
}
