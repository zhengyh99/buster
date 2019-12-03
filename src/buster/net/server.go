package net

import (
	"buster/iface"
	"buster/utils"
	"fmt"
	"net"
)

//实现IServer 接口
type Server struct {
	Name        string
	Protocal    string
	IP          string
	Port        uint32
	MsgHandler  iface.IMsgHandler            //消息处理
	ConnMng     iface.IConnManager           //链接管理
	OnConnStart func(conn iface.IConnection) //链接建立前的钩子方法
	OnConnStop  func(conn iface.IConnection) //链接建立后的钩子方法
}

func NewServer(global *utils.GlobalObj) iface.IServer {
	return &Server{
		Name:       global.ServerName,
		Protocal:   global.Protocal,
		IP:         global.IP,
		Port:       global.Port,
		MsgHandler: NewMsgHandler(),
		ConnMng:    NewConnManager(),
	}
}

//注册钩子方法：OnConnStart
func (s *Server) SetOnConnStart(hookFun func(conn iface.IConnection)) {
	s.OnConnStart = hookFun
}

//注册钩子方法：OnConnStop
func (s *Server) SetOnConnStop(hookFun func(conn iface.IConnection)) {
	s.OnConnStop = hookFun
}

//调用钩子方法：OnConnStart
func (s *Server) CallOnConnStart(conn iface.IConnection) {
	if s.OnConnStart != nil {
		fmt.Println(">>>>Call OnConnStart Function.....")
		s.OnConnStart(conn)
	}
}

//调用钩子方法：OnConnStop
func (s *Server) CallOnConnStop(conn iface.IConnection) {
	if s.OnConnStop != nil {
		fmt.Println(">>>>Call OnConnStop Function.....")
		s.OnConnStop(conn)
	}

}

//添加路由
func (s *Server) AddRouter(msgID uint32, router iface.IRouter) error {
	return s.MsgHandler.AddRouter(msgID, router)
}

//获取链接管理
func (s *Server) GetConnMng() iface.IConnManager {
	return s.ConnMng
}

//服务器开启
func (s *Server) Start() {
	fmt.Printf("服务器：%s【%s:%d】正在开启....\n", s.Name, s.IP, s.Port)
	fmt.Printf("服务器版本号：%d; 最大连接数：%d; 数据包最大值：%d \n",
		utils.GlobalObject.Version, utils.GlobalObject.MaxConn, utils.GlobalObject.MaxDataPackSize)
	go func() {
		//开启任务池
		s.MsgHandler.OpenTaskPool()
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
			//判断链接个数是否已达上限

			if s.ConnMng.Num() >= utils.GlobalObject.MaxConn {
				errInfo := NewMessage(4040, []byte("Maximum number of links has been reached"))
				dp := NewDataPack()
				errBin, _ := dp.Pack(errInfo)
				conn.Write(errBin)
				conn.Close()
				continue
			}

			//新建Connection 将conn 与客户数据CallBackToClient处理传入
			dealConn := NewConnection(s, conn, cid, s.MsgHandler)
			cid++
			go dealConn.Start()
		}
	}()

}

//服务器运行
func (s *Server) Run() {
	//开户服务
	s.Start()
	defer s.Stop()
	//TODO 服务启动后其它业务逻辑

	//阻塞状态
	select {}
}

//服务器关闭
func (s *Server) Stop() {

	fmt.Printf("Server:[%s] is stop \n", s.Name)
	s.ConnMng.ClearAll()
	//TODO 服务器的资源、状态或开辟的链接信息进行回收或停止
}
