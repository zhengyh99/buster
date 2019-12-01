package net

import (
	"buster/iface"
	"buster/utils"
	"errors"
	"fmt"
)

type MsgHandler struct {
	//存放每个msid对应该的业务处理
	MsgRouters map[uint32]iface.IRouter

	//work消息队列
	TaskQueue []chan iface.IRequest

	//Work池 work的数量
	WorkPullSize uint32
}

func NewMsgHandler() *MsgHandler {
	return &MsgHandler{
		MsgRouters:   make(map[uint32]iface.IRouter),
		WorkPullSize: utils.GlobalObject.WorkPoolSize,
		TaskQueue:    make([]chan iface.IRequest, utils.GlobalObject.WorkPoolSize),
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

//开始worker池

func (mh *MsgHandler) StartWorkPool() {
	var i uint32
	for ; i < mh.WorkPullSize; i++ {
		mh.TaskQueue[i] = make(chan iface.IRequest, utils.GlobalObject.MaxWorkTaskSize)
		go mh.StartWork(i, mh.TaskQueue[i])
	}
}

//开始单个worker

func (mh *MsgHandler) StartWork(workID uint32, taskQueue chan iface.IRequest) {
	for {

		select {
		case request := <-taskQueue:
			mh.DoMsgHandler(request)
		}
	}
}

func (mh *MsgHandler) SendToTaskWork(request iface.IRequest) {

	connID := request.GetConnection().GetConnID()
	workID := connID % mh.WorkPullSize
	fmt.Printf("Add conn id =%d, request Message id=%d to Work id =%d", connID, request.GetMsgID(), workID)

	mh.TaskQueue[workID] <- request

}
