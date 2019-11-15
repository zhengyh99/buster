package buster

import (
	"fmt"
	"net"
)

type Connection struct {
	//当前链接的socket  tcp 套接字
	Conn *net.TCPConn
	//链接ID
	ConnID uint32
	//链接状态
	isClose bool
	//当前链接处理业务的方法
	handleFun HandleFunc
	//链接退出/停止的消息通道
	ExitChan chan bool
}

//初始化链接模块的方法
func NewConnection(conn *net.TCPConn, connId uint32, callBackFun HandleFunc) *Connection {
	return &Connection{
		Conn:      conn,
		ConnID:    connId,
		handleFun: callBackFun,
		isClose:   false,
		ExitChan:  make(chan bool, 1),
	}
}

//开户链接，进入准备状态
func (c *Connection) Start() {

}

//关闭链接
func (c *Connection) Stop() {
	fmt.Printf("链接ID: %d 开始关闭", c.ConnID)
	if c.isClose == true {
		return
	}
	c.isClose = true
	close(c.ExitChan)

}

//获取当前绑定的链接
func (c *Connection) GetTCPConection() *net.TCPConn {
	return c.Conn
}

//获取当前绑定链接的ID
func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

//获取客户端状态IP port
func (c *Connection) RemoteAddr() *net.Addr {
	return c.Conn.RemoteAddr()

}

//向客户端发送数据
func (c *Connection) Send(data []byte) error {
	return nil

}
