package routers

import (
	"buster/iface"
	"buster/net"
	"fmt"
)

type DataPackRouter struct {
	net.BaseRouter
}

//处理conn 业务的Hook方法
func (dpr *DataPackRouter) Handle(r iface.IRequest) {
	fmt.Println("Call datapack Router..........")
	fmt.Printf("Message ID: %d; Data:%s", r.GetMsgID(), r.GetData())
	err := r.GetConnection().Send(1, []byte("你有新的数据请查收。。。。。"))
	if err != nil {
		fmt.Println("requst send error:", err)
	}
}
