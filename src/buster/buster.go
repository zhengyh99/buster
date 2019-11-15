package buster

import (
	"fmt"
	"net"
	"runtime"
)

//实现IServer 接口
type Server struct {
	Name     string
	Protocal string
	IP       string
	Port     string
}

func NewServer(serverName string) IServer {
	return &Server{
		Name:     serverName,
		Protocal: "tcp4",
		IP:       "127.0.0.1",
		Port:     "8808",
	}
}

//服务器开启
func (s *Server) Start() {
	fmt.Printf("服务器：%s【%s:%s】正在开启....\n", s.Name, s.IP, s.Port)
	go func() {
		//解决IP、端口
		addr, err := net.ResolveTCPAddr(s.Protocal, fmt.Sprintf("%s:%s", s.IP, s.Port))
		if err != nil {
			fmt.Println("服务器IP解析出错：", err)
		}
		//打开服务监听
		listener, err := net.ListenTCP(s.Protocal, addr)
		if err != nil {
			fmt.Println("服务器监听异常：", err)
		}
		fmt.Printf("服务器启动成功，开始监听.....")
		for {
			conn, err := listener.Accept()
			if err != nil {
				fmt.Println("服务器接收数据失败~!")
				continue
			}
			go func() {
				for {
					buf := make([]byte, 4096)
					n, err := conn.Read(buf)
					if err != nil || n <= 0 {
						fmt.Println("读取客户端数据出错：", err)
						runtime.Goexit()
						continue
					}
					fmt.Printf("客户端发来数据：%s\n", buf[:n])
					_, err = conn.Write(buf[:n])
					if err != nil {
						fmt.Println("向客户端写入数据失败：", err)
						continue
					}
				}
			}()
		}
	}()

}

//服务器运行
func (s *Server) Run() {
	//开户服务
	s.Start()
	//TODO 服务启动后其它业务逻辑

	//阻塞状态
	select {}
}

//服务器关闭
func (s *Server) Stop() {
	//TODO 服务器的资源、状态或开辟的链接信息进行回收或停止
}
