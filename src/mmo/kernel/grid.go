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
	PlayersMap map[int32]bool
	//玩家集合读写锁
	PMapLock sync.RWMutex
}

//新建格式
func NewGrid(gid, left, right, top, bottom int) *Grid {
	return &Grid{
		GID:        gid,
		Left:       left,
		Right:      right,
		Top:        top,
		Bottom:     bottom,
		PlayersMap: make(map[int32]bool),
	}
}

//添加玩家
func (g *Grid) Add(playerID int32) {
	g.PMapLock.Lock()
	defer g.PMapLock.Unlock()
	g.PlayersMap[playerID] = true
}

//删除玩家
func (g *Grid) Remove(playerID int32) {
	g.PMapLock.Lock()
	defer g.PMapLock.Unlock()
	delete(g.PlayersMap, playerID)
}

//返回所有玩家
func (g *Grid) GetPlayerIDs() (players []int32) {
	g.PMapLock.RLock()
	defer g.PMapLock.RUnlock()
	for key, _ := range g.PlayersMap {
		players = append(players, key)
	}
	return
}

//返回格式信息
func (g *Grid) String() string {
	return fmt.Sprintf("Gid:=%d,Left:%d,Right:%d,Top:%d,Bottom:%d,PlayersMap:%v",
		g.GID, g.Left, g.Right, g.Top, g.Bottom, g.PlayersMap)
}
