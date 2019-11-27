package net

import (
	"buster/iface"
	"errors"
	"fmt"
)

type MsgHandler struct {
	MsgRouters map[uint32]iface.IRouter
}

func NewMsgHandler() *MsgHandler {
	return &MsgHandler{
		MsgRouters: make(map[uint32]iface.IRouter),
	}
}

//执行对应的Router消息处理方法
func (mh *MsgHandler) DoMsgHandler(request iface.IRequest) error {
	handler, ok := mh.MsgRouters[request.GetMsgID()]
	if !ok {
		return fmt.Errorf("msg id:%d needs to be added before it can be used", request.GetMsgID())
	}

	handler.PreHandle(request)
	handler.Handle(request)
	handler.PostHandle(request)
	return nil

}

//为消息提供处理Router
func (mh *MsgHandler) AddRouter(msgID uint32, router iface.IRouter) error {
	if _, ok := mh.MsgRouters[msgID]; ok {
		return errors.New("this message id already exists")
	}
	mh.MsgRouters[msgID] = router
	return nil
}
