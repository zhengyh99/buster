package main

import (
	bnet "buster/net"
	"fmt"
	"net"
)

//封包、拆包客户端测试
func main() {
	/**
	模拟客户端
	*/

	conn, err := net.Dial("tcp", "127.0.0.1:8008")
	if err != nil {
		fmt.Println("net dial error:", err)
		return
	}
	dp := bnet.NewDataPack()
	mstr1 := []byte("hello buster!!!")
	mstr2 := []byte("你好，巴斯特！！！")

	msg1 := bnet.Message{
		ID:      1,
		DataLen: uint32(len(mstr1)),
		Data:    mstr1,
	}
	send1, err := dp.Pack(msg1)
	if err != nil {
		fmt.Println("pack msg1 error:", err)
		return
	}
	msg2 := bnet.Message{
		ID:      1,
		DataLen: uint32(len(mstr2)),
		Data:    mstr2,
	}
	send2, err := dp.Pack(msg2)
	if err != nil {
		fmt.Println("pack msg2 error:", err)
		return
	}
	send1 = append(send1, send2...)
	fmt.Println("start send")
	conn.Write(send1)

}
