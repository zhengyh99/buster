package main

import (
	"buster/bnet"
	"fmt"
	"net"
	"runtime"
)

//封包、拆包服务器端测试
func main() {
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
				dp := bnet.NewDataPack()
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
	select {}
}
