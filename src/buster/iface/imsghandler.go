package iface

/**
消息管理 抽象
*/
type IMsgHandler interface {
	//执行对应的Router消息处理方法
	DoMessageHandler(request IRequest) error
	//为消息提供处理Router
	AddRouter(msgID uint32, router IRouter) error
}
