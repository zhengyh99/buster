package buster

//定义服务器接口
type IServer interface {
	//服务器开启
	Start()
	//服务器运行
	Run()
	//服务器关闭
	Stop()
}
