package main

import (
	"buster/bnet"
	"fmt"
	"io"
	"net"
	"strconv"
	"time"
)

func main() {

	conn, err := net.Dial("tcp", ":8808")
	if err != nil {
		fmt.Println("net dial error:", err)

	}

	for i := 0; i < 10; i++ {

		dp := bnet.NewDataPack()

		binMsg, err := dp.Pack(bnet.NewMessage(uint32(i), []byte("this is busert"+strconv.Itoa(i))))
		if err != nil {
			fmt.Println("Pack error:", err)
		}
		if _, err := conn.Write(binMsg); err != nil {
			fmt.Println("conn write error:", err)
		}

		headMsg := make([]byte, dp.GetHeadLen())

		if _, err := io.ReadFull(conn, headMsg); err != nil {
			fmt.Println("read head msg error:", err)
		}

		time.Sleep(2 * time.Second)

	}

}
