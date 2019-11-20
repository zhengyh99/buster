package routers

import (
	"buster/iface"
	"fmt"
)

type PingRouter struct {
	//buster.BaseRouter
}

//处理conn 业务之前的Hook方法
func (pr *PingRouter) PreHandle(r iface.IRequest) {
	_, err := r.GetConnection().GetTCPConnection().Write([]byte("Ping...Pre ......\n"))
	if err != nil {
		fmt.Println("Ping...Pre error:", err)
	}
}

//处理conn 业务的Hook方法
func (pr *PingRouter) Handle(r iface.IRequest) {
	_, err := r.GetConnection().GetTCPConnection().Write([]byte("Ping...Ping.....Ping ......\n"))
	if err != nil {
		fmt.Println("Ping...ping error:", err)
	}
}

//处理conn 业务之后的Hook方法
func (pr *PingRouter) PostHandle(r iface.IRequest) {
	_, err := r.GetConnection().GetTCPConnection().Write([]byte("Ping...Post ......\n"))
	if err != nil {
		fmt.Println("Ping...Post error:", err)
	}
}
