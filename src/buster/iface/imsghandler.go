package iface

/**
消息管理 抽象
*/
type IMsgHandler interface {
	//执行对应的Router消息处理方法
	DoMsgHandler(request IRequest) error
	//为消息提供处理Router
	AddRouter(msgID uint32, router IRouter) error
	//开始任务池
	OpenTaskPool()
	//取模将任务分配给相应的任务队列
	SendToTask(request IRequest)
}
