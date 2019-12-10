package bnet

import (
	"buster/iface"
	"fmt"
	"sync"
)

type ConnManager struct {
	connections map[uint32]iface.IConnection
	connLock    sync.RWMutex
}

func NewConnManager() iface.IConnManager {
	return &ConnManager{
		connections: make(map[uint32]iface.IConnection),
	}
}

//添加链接
func (cm *ConnManager) Add(conn iface.IConnection) {
	cm.connLock.Lock()
	defer cm.connLock.Unlock()
	cm.connections[conn.GetConnID()] = conn
	fmt.Printf("connID:%d add to ConnManager sucessful, now number of connection is [%d]\n",
		conn.GetConnID(), cm.Num())
}

//删除链接
func (cm *ConnManager) Remove(conn iface.IConnection) {
	cm.connLock.Lock()
	defer cm.connLock.Unlock()
	delete(cm.connections, conn.GetConnID())
	conn.Stop()
	fmt.Printf("connID:%d remove from ConnManager sucessful, now number of connection is [%d]\n",
		conn.GetConnID(), cm.Num())
}

//根据id 返回链接
func (cm *ConnManager) Get(connID uint32) (iface.IConnection, error) {
	cm.connLock.RLock()
	defer cm.connLock.RUnlock()
	if conn, ok := cm.connections[connID]; ok {
		return conn, nil
	} else {
		return nil, fmt.Errorf("Connid:%d is not found!\n ", connID)
	}

}

//返回链接个数
func (cm *ConnManager) Num() int {
	return len(cm.connections)
}

//清除所有链接
func (cm *ConnManager) ClearAll() {
	cm.connLock.Lock()
	defer cm.connLock.Unlock()
	for id, conn := range cm.connections {
		//停止
		conn.Stop()
		//删除
		delete(cm.connections, id)
	}
	fmt.Printf("Clear All connection is successful, now number of connection is %d\n", cm.Num())

}
