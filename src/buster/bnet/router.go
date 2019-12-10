package bnet

import "buster/iface"

type BaseRouter struct{}

//处理conn 业务之前的Hook方法
func (br *BaseRouter) PreHandle(r iface.IRequest) {}

//处理conn 业务的Hook方法
func (br *BaseRouter) Handle(r iface.IRequest) {}

//处理conn 业务之后的Hook方法
func (br *BaseRouter) PostHandle(r iface.IRequest) {}
