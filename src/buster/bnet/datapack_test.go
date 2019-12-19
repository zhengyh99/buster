package bnet

import (
	"fmt"
	"net"
	"runtime"
	"testing"
)

func TestDataPack(t *testing.T) {
	/**
	服务器端
	*/
	//打开监听
	addr, err := net.ResolveTCPAddr("tcp4", fmt.Sprintf("%s:%d", "127.0.0.1", 8008))
	if err != nil {
		fmt.Println("服务器IP解析出错：", err)
	}
	listener, err := net.ListenTCP("tcp4", addr)
	if err != nil {
		fmt.Println("net listen error:", err)
	}
	go func() {
		for {
			//接收数据
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("listen accept error:", err)
				runtime.Goexit()
			}
			go func(conn *net.TCPConn) {
				dp := NewDataPack()
				for {
					msg, err := dp.UnPack(conn)
					if err != nil {
						fmt.Println("error:", err)
					}
					fmt.Printf("\t Resv Message ID:%d,Len:%d,Data:%s\n\n", msg.GetID(),
						msg.GetDataLen(), msg.GetData())
				}

			}(conn)
		}

	}()

	/**
	模拟客户端
	*/

	conn, err := net.Dial("tcp", "127.0.0.1:8008")
	if err != nil {
		fmt.Println("net dial error:", err)
		return
	}
	dp := NewDataPack()
	mstr1 := []byte("hello buster!!!")
	mstr2 := []byte("你好，巴斯特！！！")

	msg1 := &Message{
		ID:      1,
		DataLen: uint32(len(mstr1)),
		Data:    mstr1,
	}
	send1, err := dp.Pack(msg1)
	if err != nil {
		fmt.Println("pack msg1 error:", err)
		return
	}
	msg2 := &Message{
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
	conn.Write(send1)
	select {}

}
