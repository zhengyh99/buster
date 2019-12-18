package bnet

import (
	"buster/iface"
	"buster/utils"
	"errors"
	"fmt"
	"io"
	"net"
	"sync"
)

type Connection struct {

	//当前链接隶属服务器
	MyServer iface.IServer
	//当前链接的socket  tcp 套接字
	Conn *net.TCPConn
	//链接ID
	ConnID uint32
	//链接状态
	isClose bool

	//客户端读写消息通道
	MsgChan chan []byte

	//链接退出/停止的消息通道
	ExitChan chan bool
	//当前链接对应的router业务处理
	MsgHandler iface.IMsgHandler

	//链接属性
	property map[string]interface{}
	//链接属性操作锁
	propertyLock sync.RWMutex
}

//初始化链接模块的方法
func NewConnection(server iface.IServer, conn *net.TCPConn, connId uint32, msgHandler iface.IMsgHandler) *Connection {
	c := &Connection{
		MyServer:   server,
		Conn:       conn,
		ConnID:     connId,
		isClose:    false,
		MsgHandler: msgHandler,
		MsgChan:    make(chan []byte),
		ExitChan:   make(chan bool, 1),
		property:   make(map[string]interface{}),
	}
	//添加链接至链接管理中
	server.GetConnMng().Add(c)
	return c
}

//向客户端发送消息
func (c *Connection) StartWrite() {
	fmt.Printf("Goroutine 链接ID:%d 准备向客户端发送数据\n", c.ConnID)
	for {
		select { //监听消息通道
		//接收读消息
		case data := <-c.MsgChan:

			c.Conn.Write(data)
		//接收退出消息
		case <-c.ExitChan:
			return
		}
	}
}

//开始读取客户端消息
func (c *Connection) StartRead() {
	fmt.Printf("Goroutine 链接ID:%d 开始接收客户端数据\n", c.ConnID)
	defer fmt.Printf("链接ID:%d 正在退出 ，远程IP地址为：%s\n", c.ConnID, c.RemoteAddr().String())
	defer c.Stop()
	for {
		//对接收的数据拆包

		msg, err := c.getMsg()
		if err != nil {
			fmt.Printf("connection get Message error :", err)
			break
		}

		//定义一个Resuest
		req := &Request{
			conn: c,
			msg:  msg,
		}
		//判断是否打开任务池机制
		if utils.GlobalConfig.MaxTaskSize > 0 {
			//将requst发送给任务池处理
			c.MsgHandler.SendToTask(req)
		} else {
			//从路由中找到conn对应该的router
			//从msgid找到对应的业务处理
			go c.MsgHandler.DoMsgHandler(req)
		}

	}

}

//开户链接，进入准备状态
func (c *Connection) Start() {
	fmt.Printf("远程IP地址为：%s, 链接ID:%d 正在运行\n", c.RemoteAddr().String(), c.ConnID)
	go c.StartRead()
	go c.StartWrite()
	//调用链接开启的钩子方法
	c.MyServer.CallOnConnStart(c)
}

//关闭链接
func (c *Connection) Stop() {
	fmt.Printf("链接ID: %d 开始关闭\n", c.ConnID)
	if c.isClose == true {
		return
	}
	c.isClose = true
	c.ExitChan <- true
	//调用链接关闭钩子方法
	c.MyServer.CallOnConnStop(c)
	//关闭链接
	c.Conn.Close()
	//将链接从链接管理中移除
	c.MyServer.GetConnMng().Remove(c)
	//回收资源

	close(c.ExitChan)
	close(c.MsgChan)

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
func (c *Connection) Send(msgID uint32, msgData []byte) error {
	if c.isClose {
		return errors.New(" connection is closed ")
	}
	//建立消息
	msg := NewMessage(msgID, msgData)
	//封包
	dp := NewDataPack()
	binMsg, err := dp.Pack(msg)
	if err != nil {
		return err
	}
	//通过消息通道发送给 客户端
	c.MsgChan <- binMsg

	return nil

}

//设置属性
func (c *Connection) SetProperty(key string, value interface{}) {
	c.propertyLock.Lock()
	defer c.propertyLock.Unlock()
	c.property[key] = value
}

//获取属性
func (c *Connection) GetProperty(key string) interface{} {
	c.propertyLock.RLock()
	defer c.propertyLock.RUnlock()
	if v, ok := c.property[key]; ok {
		return v
	}
	fmt.Printf(" property:%s if not found.\n ", key)
	return nil

}

//删除属性
func (c *Connection) RemoveProperty(key string) {
	c.propertyLock.Lock()
	defer c.propertyLock.Unlock()
	delete(c.property, key)
}

func (c *Connection) getMsg() (iface.IMessage, error) {
	dp := NewDataPack()
	headData := make([]byte, dp.GetHeadLen())
	if _, err := io.ReadFull(c.GetTCPConnection(), headData); err != nil {
		fmt.Println("read Message head error:", err)
		return nil, err
	}
	msg, err := dp.UnPack(headData)
	if err != nil {
		fmt.Println("datapack error:", err)
		return nil, err
	}
	var data []byte
	if msg.GetDataLen() > 0 {
		data = make([]byte, msg.GetDataLen())
		if _, err := io.ReadFull(c.GetTCPConnection(), data); err != nil {
			fmt.Println("read Message head error:", err)
			return nil, err
		}
	}
	msg.SetData(data)
	return msg, nil

}
