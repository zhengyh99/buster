package bnet

import (
	"fmt"
	"io"
	"net"
	"runtime"
	"testing"
)

func TestDataPack(t *testing.T) {
	/**
	服务器端
	*/
	//打开监听
	listener, err := net.Listen("tcp4", "127.0.0.1:8008")
	if err != nil {
		fmt.Println("net listen error:", err)
	}
	go func() {
		for {
			//接收数据
			conn, err := listener.Accept()
			if err != nil {
				fmt.Println("listen accept error:", err)
				runtime.Goexit()
			}
			go func(conn net.Conn) {
				dp := NewDataPack()
				for {
					headData := make([]byte, dp.GetHeadLen())
					_, err := io.ReadFull(conn, headData)
					if err != nil {
						fmt.Println(" read headata error:", err)
						break
					}
					msgHead, err := dp.UnPack(headData)
					if err != nil {
						fmt.Println(" dp UnPack error:", err)
						return
					}
					if msgHead.GetDataLen() > 0 {
						msg, ok := msgHead.(*Message)
						if ok {
							msg.Data = make([]byte, msg.GetDataLen())
							_, err := io.ReadFull(conn, msg.Data)
							if err != nil {
								fmt.Println("read msg data error:", err)
								return
							}
						} else {
							fmt.Println("消息类型转换错误")
							return
						}
						fmt.Printf("\t Resv Message ID:%d,Len:%d,Data:%s\n\n", msg.GetID(),
							msg.GetDataLen(), msg.GetData())
					}
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

	msg1 := Message{
		ID:      1,
		DataLen: uint32(len(mstr1)),
		Data:    mstr1,
	}
	send1, err := dp.Pack(msg1)
	if err != nil {
		fmt.Println("pack msg1 error:", err)
		return
	}
	msg2 := Message{
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
