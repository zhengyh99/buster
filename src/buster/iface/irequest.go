package iface

/**
IRequest接口 实现对客户端链接和请求数据封装
*/
type IRequest interface {
	//得到当前链接
	GetConnection() IConnection
	//得到请求数据
	GetData() []byte

	GetMsgID() uint32
}
