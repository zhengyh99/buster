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
	Send(msgID uint32, msgData []byte) error
	
	//设置属性
	SetProperty(key string, value interface{})
	//获取属性
	GetProperty(key string) (value interface{}, err error)
	//删除属性
	RemoveProperty(key string) 

//定义处理链接业务的方法
type HandleAPI func(*net.TCPConn, []byte, int) error
