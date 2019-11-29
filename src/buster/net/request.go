package net

import "buster/iface"

type Request struct {
	conn iface.IConnection
	msg  iface.IMessage
}

//得到当前链接
func (r *Request) GetConnection() iface.IConnection {
	return r.conn
}

//得到请求数据
func (r *Request) GetData() []byte {
	return r.msg.GetData()

}

//获取 message ID
func (r *Request) GetMsgID() uint32 {
	return r.msg.GetID()
}

//向客户端发送数据
func (r *Request) Send(msgID uint32, data []byte) error {
	return r.conn.Send(msgID, data)
}
