package iface

type IRouter interface {
	//处理conn 业务之前的Hook方法
	PreHandle(r IRequest)
	//处理conn 业务的Hook方法
	Handle(r IRequest)
	//处理conn 业务之后的Hook方法
	PostHandle(r IRequest)
}
