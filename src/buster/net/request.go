package net

import "buster/iface"

type Request struct {
	conn iface.IConnection
	data []byte
}

//得到当前链接
func (r *Request) GetConnection() iface.IConnection {
	return r.conn
}

//得到请求数据
func (r *Request) GetData() []byte {
	return r.data

}
