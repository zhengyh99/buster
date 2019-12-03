package iface

//定义服务器接口
type IServer interface {
	//服务器开启
	Start()
	//服务器运行
	Run()
	//服务器关闭
	Stop()
	//添加路由
	AddRouter(msgID uint32, router IRouter) error

	//获取链接管理
	GetConnMng() IConnManager
	//注册钩子方法：OnConnStart
	SetOnConnStart(hookFun func(conn IConnection))
	//注册钩子方法：OnConnStop
	SetOnConnStop(hookFun func(conn IConnection))

	//调用钩子方法：OnConnStart
	CallOnConnStart(conn IConnection)
	//调用钩子方法：OnConnStop
	CallOnConnStop(conn IConnection)
}
