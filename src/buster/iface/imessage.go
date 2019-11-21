package iface

type IMessage interface {

	//获取消息ID
	GetID() uint32
	//设置消息ID
	SetID(id uint32)

	//获取消息长度
	GetDataLen() uint32
	//设置消息长度
	SetDatalen(len uint32)

	//获取消息内容
	GetData() []byte
	//设置消息内容
	SetData(data []byte)
}
