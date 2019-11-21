package net

import (
	"buster/iface"
	"buster/utils"
	"bytes"
	"encoding/binary"
	"errors"
)

type DataPack struct{}

//获取包头长度{
func (dp *DataPack) GetHeadLen() uint32 {
	return 8
}

//封包
func (dp *DataPack) Pack(msg iface.IMessage) (data []byte, err error) {
	//创建一个bf ：存放bytes的缓冲
	bf := bytes.NewBuffer([]byte{})
	//将datalen 写入缓冲
	if err = binary.Write(bf, binary.LittleEndian, msg.GetDataLen()); err != nil {
		return nil, err
	}
	//将id 写入缓冲
	if err = binary.Write(bf, binary.LittleEndian, msg.GetID()); err != nil {
		return nil, err
	}
	//将 data 写入缓冲
	if err = binary.Write(bf, binary.LittleEndian, msg.GetDataLen()); err != nil {
		return nil, err
	}
	return bf.Bytes(), err
}

//拆包
func (dp *DataPack) UnPack(data []byte) (iface.IMessage, error) {
	msg := &Message{}
	bf := bytes.NewBuffer([]byte{})

	if err := binary.Read(bf, binary.LittleEndian, msg.DataLen); err != nil {
		return nil, err
	}
	if err := binary.Read(bf, binary.LittleEndian, msg.ID); err != nil {
		return nil, err
	}

	if utils.GlobalObject.MaxDataPackSize > 0 && utils.GlobalObject.MaxDataPackSize < msg.DataLen {
		return nil, errors.New("Data length exceeds limit")
	}
	return msg, nil

}
