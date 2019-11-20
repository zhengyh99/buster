package net

import (
	"buster/iface"
	"buster/utils"
	"fmt"
	"net"
	"runtime"
)

type Connection struct {
	//当前链接的socket  tcp 套接字
	Conn *net.TCPConn
	//链接ID
	ConnID uint32
	//链接状态
	isClose bool

	//链接退出/停止的消息通道
	ExitChan chan bool
	//当前链接对应的router业务处理
	Router iface.IRouter
}

//初始化链接模块的方法
func NewConnection(conn *net.TCPConn, connId uint32, router iface.IRouter) *Connection {
	return &Connection{
		Conn:     conn,
		ConnID:   connId,
		isClose:  false,
		Router:   router,
		ExitChan: make(chan bool, 1),
	}
}
func (c *Connection) StartRead() {
	fmt.Printf("Goroutine 链接ID:%d 开始运行\n", c.ConnID)
	defer fmt.Printf("链接ID:%d 正在退出 ，远程IP地址为：%s\n", c.ConnID, c.RemoteAddr().String())
	defer c.Stop()
	for {
		buf := make([]byte, utils.GlobalObject.MaxDataPackSize)
		n, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("conn read error:", err)
			runtime.Goexit()
		}
		//定义一个Resuest
		req := &Request{
			conn: c,
			data: buf[:n],
		}
		c.Router.PreHandle(req)
		c.Router.Handle(req)
		c.Router.PostHandle(req)

	}

}

//开户链接，进入准备状态
func (c *Connection) Start() {
	fmt.Printf("远程IP地址为：%s, 链接ID:%d 正在运行\n", c.RemoteAddr().String(), c.ConnID)
	go c.StartRead()
}

//关闭链接
func (c *Connection) Stop() {
	fmt.Printf("链接ID: %d 开始关闭\n", c.ConnID)
	if c.isClose == true {
		return
	}
	c.isClose = true
	close(c.ExitChan)

}

//获取当前绑定的链接
func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

//获取当前绑定链接的ID
func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

//获取客户端状态IP port
func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()

}

//向客户端发送数据
func (c *Connection) Send(data []byte) error {
	return nil

}
