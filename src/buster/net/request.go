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

func (r *Request) GetMsgID() uint32 {
	return r.msg.GetID()
}
