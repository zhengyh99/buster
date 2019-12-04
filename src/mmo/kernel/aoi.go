package kernel

import "fmt"

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

func NewAOIManager(left, right, gridsX, top, bottom, gridsY int) *AOIManager {
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

// 重写 String()
func (am *AOIManager) String() string {
	s := fmt.Sprintf("AOIManage:\nLeft:%d,Right:%d,GridsX:%d,Top:%d,Bottom:%d,GridsY:%d\n",
		am.Left, am.Right, am.GridsX, am.Top, am.Bottom, am.GridsY)
	for k, v := range am.GridMap {
		s += fmt.Sprintf("Grip:%d \n %s \n", k, v)
	}
	return s
}
