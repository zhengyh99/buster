package routers

import (
	"buster/bnet"
	"buster/iface"
	"fmt"
)

type DataPackRouter struct {
	bnet.BaseRouter
}

//处理conn 业务的Hook方法
func (dpr *DataPackRouter) Handle(r iface.IRequest) {
	fmt.Println("Call datapack Router..........")
	fmt.Printf("Message ID: %d; Data:%s\n", r.GetMsgID(), r.GetData())
	err := r.Send(1, []byte("你有新的数据请查收。。。。。"))
	if err != nil {
		fmt.Println("requst send error:", err)
	}
}
