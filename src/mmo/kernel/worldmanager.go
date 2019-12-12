package kernel

import "sync"

//定义世界管理
type WorldManager struct {
	AoiMng  *AOIManager
	Players map[int32]*Player
	PLock   sync.RWMutex
}

//定义全局玩家世界
var WorldMng *WorldManager

//初始化玩家世界
func init() {
	WorldMng = &WorldManager{
		AoiMng:  NewAOIManager(AOILeft, AOIRight, AOITop, AOIBottom, AOIGridsX, AOIGridsY),
		Players: make(map[int32]*Player),
	}
}

//向世界管理中添加玩家
func (wm *WorldManager) AddPlayer(player *Player) {
	wm.PLock.Lock()
	defer wm.PLock.Unlock()
	wm.Players[player.Pid] = player

	wm.AoiMng.AddPlayerToPosition(player.Pid, player.X, player.Y)
}

//从世界中移出玩家
func (wm *WorldManager) RemovePlayer(pid int32) {
	wm.PLock.Lock()
	defer wm.PLock.Unlock()
	player := wm.Players[pid]
	if player != nil {
		wm.AoiMng.RemovePlayerFromPosition(player.Pid, player.X, player.Y)
	}
}

//通过pid 返回玩家世界中的玩家
func (wm *WorldManager) GetPlayer(pid int32) *Player {
	wm.PLock.RLock()
	defer wm.PLock.RUnlock()
	return wm.Players[pid]
}

//返回玩家世界中的所有玩家
func (wm *WorldManager) GetAllPlayer() (players []*Player) {
	players = make([]*Player, 0)
	for _, player := range wm.Players {
		players = append(players, player)
	}
	return
}
