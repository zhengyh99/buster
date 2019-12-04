package kernel

import (
	"fmt"
	"testing"
)

func TestAOIManager(t *testing.T) {
	//初始化AOIManager
	am := NewAOIManager(100, 300, 4, 200, 350, 5)
	//打印
	fmt.Println(am)
}
