package main

import (
	"buster/bnet"
	"buster/iface"
	"buster/routers"
	"fmt"
	"strconv"
	"time"
)

func OnConnBegin(conn iface.IConnection) {
	fmt.Println("/////// on conn begin is calling.....")
	conn.Send(202, []byte("Connection is established "))
	conn.SetProperty("cname", "链接："+strconv.Itoa(int(conn.GetConnID())))
	conn.SetProperty("QQ", 25523051)
	conn.SetProperty("Date", time.Now())

}

func OnConnLost(conn iface.IConnection) {
	fmt.Println("///////// on conn lost is calling")
	fmt.Printf("conn id :%d is lost........\n", conn.GetConnID())
	cname := conn.GetProperty("cname")
	QQ := conn.GetProperty("QQ")
	Date := conn.GetProperty("Date").(time.Time)
	layout := "2006年1月2日 15 ：04：05"
	fmt.Printf("CName:%s,\nQQ:%d\n,Date:%s\n\n", cname, QQ, Date.Format(layout))
}
func main() {
	server := bnet.NewServer()
	server.SetOnConnStart(OnConnBegin)
	server.SetOnConnStop(OnConnLost)
	server.AddRouter(0, &routers.DataPackRouter{})
	server.AddRouter(1, &routers.PingRouter{})
	server.Run()
}
