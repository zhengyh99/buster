package iface

type IConnManager interface {
	//添加链接
	Add(conn IConnection)
	//删除链接
	Remove(conn IConnection)
	//根据id 返回链接
	Get(connID uint32) (IConnection, error)
	//返回链接个数
	Num() int
	//清除所有链接
	ClearAll()
}
