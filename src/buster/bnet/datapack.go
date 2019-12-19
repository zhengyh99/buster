package bnet

import (
	"buster/iface"
	"buster/utils"
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"net"
)

type DataPack struct{}

func NewDataPack() *DataPack {
	return &DataPack{}
}

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
	if err = binary.Write(bf, binary.LittleEndian, msg.GetData()); err != nil {
		return nil, err
	}
	return bf.Bytes(), err
}

//拆包
func (dp *DataPack) UnPack(conn net.Conn) (iface.IMessage, error) {

	//获取并解析包头
	headData := make([]byte, dp.GetHeadLen())
	if _, err := io.ReadFull(conn, headData); err != nil {
		fmt.Println("read Message head error:", err)
		return nil, err
	}
	msg := &Message{}
	bf := bytes.NewBuffer(headData)
	//读长度msg.DataLen
	if err := binary.Read(bf, binary.LittleEndian, &msg.DataLen); err != nil {
		fmt.Println("binary read1 error:", err)
		return nil, err
	}

	//读id msg.ID
	if err := binary.Read(bf, binary.LittleEndian, &msg.ID); err != nil {
		fmt.Println("binary read2 error:", err)
		return nil, err
	}
	//包容量限制
	if utils.GlobalConfig.MaxDataPackSize > 0 && utils.GlobalConfig.MaxDataPackSize < msg.DataLen {
		return nil, errors.New("Data length exceeds limit")
	}

	//读消息体 msg.Data
	var data []byte
	if msg.GetDataLen() > 0 {
		data = make([]byte, msg.GetDataLen())
		if _, err := io.ReadFull(conn, data); err != nil {
			fmt.Println("read Message head error:", err)
			return nil, err
		}
	}
	msg.SetData(data)

	fmt.Println(">>>>>>>>,,,,, msg", msg)
	return msg, nil

}
