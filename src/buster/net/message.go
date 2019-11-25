package net

import (
	"buster/iface"
)

type Message struct {
	ID      uint32
	DataLen uint32
	Data    []byte
}

func NewMessage(id uint32, data []byte) iface.IMessage {
	return &Message{
		ID:      id,
		DataLen: uint32(len(data)),
		Data:    data,
	}
}

//获取消息ID
func (m Message) GetID() uint32 {
	return m.ID
}

//设置消息ID
func (m Message) SetID(id uint32) {
	m.ID = id
}

//获取消息长度
func (m Message) GetDataLen() uint32 {
	return m.DataLen
}

//设置消息长度
func (m Message) SetDatalen(len uint32) {
	m.DataLen = len
}

//获取消息内容
func (m Message) GetData() []byte {
	return m.Data
}

//设置消息内容
func (m Message) SetData(data []byte) {
	m.Data = data
}
