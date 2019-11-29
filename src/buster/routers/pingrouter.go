package routers

import (
	"buster/iface"
	"buster/net"
	"fmt"
)

type PingRouter struct {
	net.BaseRouter
}

//处理conn 业务之前的Hook方法
func (pr *PingRouter) PreHandle(r iface.IRequest) {
	fmt.Println("Call pre ping Router..........")
	fmt.Printf("Message ID: %d; Data:%s\n", r.GetMsgID(), r.GetData())
	err := r.Send(1, []byte("pre	pre 	Pinging。。。。。"))
	if err != nil {
		fmt.Println("requst send error:", err)
	}
}

//处理conn 业务的Hook方法
func (pr *PingRouter) Handle(r iface.IRequest) {
	fmt.Println("Call pingping Router..........")
	fmt.Printf("Message ID: %d; Data:%s\n", r.GetMsgID(), r.GetData())
	err := r.Send(1, []byte("Pinging	Pinging			Pinging。。。。。"))
	if err != nil {
		fmt.Println("requst send error:", err)
	}
}

//处理conn 业务之后的Hook方法
func (pr *PingRouter) PostHandle(r iface.IRequest) {
	fmt.Println("Call post pingping Router..........")
	fmt.Printf("Message ID: %d; Data:%s\n", r.GetMsgID(), r.GetData())
	err := r.Send(1, []byte("Pinging	posting 	posting 。。。。。"))
	if err != nil {
		fmt.Println("requst send error:", err)
	}
}
