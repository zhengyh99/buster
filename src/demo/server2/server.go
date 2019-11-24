package main

import (
	bnet "buster/net"
	"fmt"
	"io"
	"net"
	"runtime"
)

func main() {
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
				dp := bnet.NewDataPack()
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
						msg, ok := msgHead.(*bnet.Message)
						fmt.Println("OK:", ok)
						if ok {
							msg.Data = make([]byte, msg.GetDataLen())
							_, err := io.ReadFull(conn, msg.Data)
							fmt.Printf(" read start....%s\n", msg.Data)
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
	select {}
}
