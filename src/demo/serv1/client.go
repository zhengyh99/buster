package main

import (
	"fmt"
	"net"
	"time"
)

func main() {

	conn, err := net.Dial("tcp", ":8808")
	if err != nil {
		fmt.Println("net dial error:", err)

	}

	for i := 0; i < 10; i++ {
		_, err := conn.Write([]byte("Hello Buster!!!"))
		if err != nil {
			fmt.Println("conn write error:", err)
			break
		}
		buf := make([]byte, 4096)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("conn read error:", err)
			break
		}
		fmt.Printf("Server Call Back Message：\n%s\n\n", string(buf[:n]))
		time.Sleep(time.Second * 2)

	}

}
