package iface

import "net"

//定义tcp 链接接口
type IConnection interface {
	//开户链接，进入准备状态
	Start()
	//关闭链接
	Stop()
	//获取当前绑定的链接
	GetTCPConnection() *net.TCPConn
	//获取当前绑定链接的ID
	GetConnID() uint32
	//获取客户端状态IP port
	RemoteAddr() net.Addr
	//向客户端发送数据
	Send(data []byte) error
}

//定义处理链接业务的方法
type HandleAPI func(*net.TCPConn, []byte, int) error
