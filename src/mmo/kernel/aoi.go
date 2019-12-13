package kernel

import "fmt"

const (
	AOILeft   = 110
	AOIRight  = 310
	AOITop    = 10
	AOIBottom = 110
	AOIGridsX = 20
	AOIGridsY = 10
)

type AOIManager struct {
	//区域边界 左
	Left int
	//区域边界 右
	Right int
	//X轴方向 grid 个数
	GridsX int

	//区域边界 上
	Top int
	//区域边界 下
	Bottom int
	//y轴方向 grid 个数
	GridsY int

	//区域内 Grid布局
	GridMap map[int]*Grid
}

func NewAOIManager(left, right, top, bottom, gridsX, gridsY int) *AOIManager {
	am := &AOIManager{
		Left:    left,
		Right:   right,
		GridsX:  gridsX,
		Top:     top,
		Bottom:  bottom,
		GridsY:  gridsY,
		GridMap: make(map[int]*Grid),
	}
	//初始化 grid 布局
	for y := 0; y < am.GridsY; y++ {
		for x := 0; x < am.GridsX; x++ {

			gid := y*am.GridsY + x
			fmt.Println("gid:", gid)
			am.GridMap[gid] = NewGrid(gid,
				am.Left+x*am.gridWidth(),
				am.Left+(x+1)*am.gridWidth(),
				am.Top+y*am.gridHeight(),
				am.Top+(y+1)*am.gridHeight())
		}
	}

	return am

}

//返回 grid 宽度
func (am *AOIManager) gridWidth() int {
	return (am.Right - am.Left) / am.GridsX
}

//返回 grid 高度
func (am *AOIManager) gridHeight() int {
	return (am.Bottom - am.Top) / am.GridsY
}

//根据gid 返回当前Grid 周围环境（九宫格）格子集合
func (am *AOIManager) GetSurroundingsByGid(gid int) (grids []*Grid) {
	grids = make([]*Grid, 0)
	//判断gid 是否存在
	if _, ok := am.GridMap[gid]; !ok {
		return nil
	}
	//将当前格子 加入九宫格集合
	grids = append(grids, am.GridMap[gid])

	//当前格式在X轴上的序号（从0开始）
	gX := gid % am.GridsX
	//判断左侧是否有格子，如果有，加入集合
	if gX > 0 {
		grids = append(grids, am.GridMap[gid-1])
	}
	//判断右侧是否有格子，如果有，加入集合
	if gX < am.GridsX-1 {
		grids = append(grids, am.GridMap[gid+1])
	}

	//将X轴方向所有格子 入Xgrids集合中
	Xgrids := make([]int, 0, len(grids))
	for _, v := range grids {
		Xgrids = append(Xgrids, v.GID)
	}

	//遍历X轴格式，判断每个格子的上下是否有格式，如果有加入集合
	for _, v := range Xgrids {
		//当前格式在Y轴上的序号（从0开始）
		gY := gid / am.GridsX
		//判断上方是否有格式
		if gY > 0 {
			grids = append(grids, am.GridMap[v-am.GridsX])
		}
		//判断下方是否有格子
		if gY < am.GridsY-1 {
			grids = append(grids, am.GridMap[v+am.GridsX])
		}
	}
	return
}

//根据坐标获取 坐标对应格子id
func (am *AOIManager) GetGidByPosition(x, y float32) (gid int) {
	//X轴方向对应该格子序数
	gridX := int(x) - am.Left/am.gridWidth()
	//Y轴方向对应该格子序数
	gridY := int(y) - am.Top/am.gridHeight()
	gid = gridX/am.GridsX + gridY
	return
}

//根据坐标返回周边格子的玩家ID
func (am *AOIManager) GetPlayersByPosition(x, y float32) (playerIDs []int32) {
	//根据坐标返回所在格子的ID
	gid := am.GetGidByPosition(x, y)
	fmt.Println("gid:--------1:", gid)
	//根据格子ID返回周围格子集合
	grids := am.GetSurroundingsByGid(gid)
	fmt.Println("neighbor gid :", grids)

	for _, grid := range grids {
		playerIDs = append(playerIDs, grid.GetPlayerIDs()...)

	}
	return
}

// 重写 String()
func (am *AOIManager) String() string {
	s := fmt.Sprintf("AOIManage:\nLeft:%d,Right:%d,GridsX:%d,Top:%d,Bottom:%d,GridsY:%d\n",
		am.Left, am.Right, am.GridsX, am.Top, am.Bottom, am.GridsY)
	for k, v := range am.GridMap {
		s += fmt.Sprintf("Grip:%d \n %s \n", k, v)
	}
	return s
}

//向Grid:gid中 加入Player: pid
func (am *AOIManager) AddPlayerToGrid(pid int32, gid int) {
	am.GridMap[gid].Add(pid)
}

//将Grid:gid中Player: pid 删除
func (am *AOIManager) RemovePlayerFromGrid(pid int32, gid int) {
	am.GridMap[gid].Remove(pid)
}

//从Grid:gid中 获取所有Player的pid
func (am *AOIManager) GetPlayerIDsFromGrid(gid int) (playerIDs []int32) {
	return am.GridMap[gid].GetPlayerIDs()
}

//向坐标系 （x,y）对应的Grid:gid中加入Player:pid
func (am *AOIManager) AddPlayerToPosition(pid int32, x, y float32) {
	gid := am.GetGidByPosition(x, y)
	am.AddPlayerToGrid(pid, gid)

}

func (am *AOIManager) RemovePlayerFromPosition(pid int32, x, y float32) {
	gid := am.GetGidByPosition(x, y)
	am.RemovePlayerFromGrid(pid, gid)
}
