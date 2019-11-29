package net

import (
	"buster/iface"
	"buster/utils"
	"fmt"
	"net"
)

//实现IServer 接口
type Server struct {
	Name       string
	Protocal   string
	IP         string
	Port       uint32
	MsgHandler iface.IMsgHandler
}

func NewServer(global *utils.GlobalObj) iface.IServer {
	return &Server{
		Name:       global.ServerName,
		Protocal:   global.Protocal,
		IP:         global.IP,
		Port:       global.Port,
		MsgHandler: NewMsgHandler(),
	}
}

//添加路由
func (s *Server) AddRouter(msgID uint32, router iface.IRouter) error {
	return s.MsgHandler.AddRouter(msgID, router)
}

//服务器开启
func (s *Server) Start() {
	fmt.Printf("服务器：%s【%s:%d】正在开启....\n", s.Name, s.IP, s.Port)
	go func() {
		//解决IP、端口
		addr, err := net.ResolveTCPAddr(s.Protocal, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("服务器IP解析出错：", err)
		}
		//打开服务监听
		listener, err := net.ListenTCP(s.Protocal, addr)
		if err != nil {
			fmt.Println("服务器监听异常：", err)
		}
		fmt.Printf("服务器启动成功，开始监听.....")
		var cid uint32
		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("服务器接收数据失败~!")
				continue
			}
			//新建Connection 将conn 与客户数据CallBackToClient处理传入
			dealConn := NewConnection(conn, cid, s.MsgHandler)
			cid++
			go dealConn.Start()
		}
	}()

}

//服务器运行
func (s *Server) Run() {
	//开户服务
	s.Start()
	//TODO 服务启动后其它业务逻辑

	//阻塞状态
	select {}
}

//服务器关闭
func (s *Server) Stop() {
	//TODO 服务器的资源、状态或开辟的链接信息进行回收或停止
}
