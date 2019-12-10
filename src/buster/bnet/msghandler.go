package bnet

import (
	"buster/iface"
	"buster/utils"
	"errors"
	"fmt"
)

type MsgHandler struct {
	//存放每个msid对应该的业务处理
	MsgRouters map[uint32]iface.IRouter

	//任务队列
	TaskQueue []chan iface.IRequest

	//任务池 中任务的数量
	TaskPullSize uint32
}

func NewMsgHandler() *MsgHandler {
	return &MsgHandler{
		MsgRouters:   make(map[uint32]iface.IRouter),
		TaskPullSize: utils.GlobalConfig.TaskPoolSize,
		TaskQueue:    make([]chan iface.IRequest, utils.GlobalConfig.TaskPoolSize),
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

//开始任务池
func (mh *MsgHandler) OpenTaskPool() {
	var i uint32
	for ; i < mh.TaskPullSize; i++ {
		mh.TaskQueue[i] = make(chan iface.IRequest, utils.GlobalConfig.MaxTaskSize)
		go mh.RunTask(i, mh.TaskQueue[i])
	}
}

//开始单个任务
func (mh *MsgHandler) RunTask(taskID uint32, taskQueue chan iface.IRequest) {
	fmt.Printf("Task ID: %d is runing.....\n", taskID)
	for {

		select {
		case request := <-taskQueue:
			mh.DoMsgHandler(request)
		}
	}
}

//取模将任务分配给相应的任务队列
func (mh *MsgHandler) SendToTask(request iface.IRequest) {

	connID := request.GetConnection().GetConnID()
	taskID := connID % mh.TaskPullSize
	fmt.Printf("Add conn id =%d, request Message id=%d to task id =%d", connID, request.GetMsgID(), taskID)

	mh.TaskQueue[taskID] <- request

}
