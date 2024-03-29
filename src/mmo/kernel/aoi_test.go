package kernel

import (
	"fmt"
	"testing"
)

func TestAOIManager(t *testing.T) {
	//初始化AOIManager
	am := NewAOIManager(AOILeft, AOIRight, AOITop, AOIBottom, AOIGridsX, AOIGridsY)
	//打印
	fmt.Println(am)
}

func TestGetSurroundingsByGid(t *testing.T) {
	//初始化AOIManager
	am := NewAOIManager(AOILeft, AOIRight, AOITop, AOIBottom, AOIGridsX, AOIGridsY)

	for gid, _ := range am.GridMap {

		grids := am.GetSurroundingsByGid(gid)
		fmt.Println("gid:", gid, "grid Len", len(grids))
		allGid := make([]int, 0, len(grids))
		for _, grid := range grids {
			allGid = append(allGid, grid.GID)
		}
		fmt.Println("周围grid有：", allGid)

	}
}
