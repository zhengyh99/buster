package utils

import (
	"buster/iface"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type GlobalObj struct {
	Server          iface.IServer
	IP              string
	Port            uint32
	ServerName      string
	Protocal        string
	Version         string
	MaxConn         int
	MaxDataPackSize uint32
	WorkPoolSize    uint32 //work池中work(Goruntine)数量
	MaxWorkTaskSize uint32 //允许用户开辟work的最大数量
}

func (gObj *GlobalObj) Reload() {
	data, err := ioutil.ReadFile("config/conf.json")
	if err != nil {
		fmt.Println("read config file error:", err)
	}
	err = json.Unmarshal(data, gObj)
	if err != nil {
		fmt.Println("Config file unmarshal error:", err)
	} else {
		fmt.Println("Config Json 文件读取成功。。。。")
	}

}

var GlobalObject *GlobalObj

func init() {
	GlobalObject = &GlobalObj{
		//以下为默认配置
		IP:              "0.0.0.0",
		Port:            8808,
		ServerName:      "Buster Server",
		Protocal:        "tcp4",
		Version:         "1.0",
		MaxConn:         3,
		MaxDataPackSize: 4096,
		WorkPoolSize:    10,
		MaxWorkTaskSize: 1024,
	}
	GlobalObject.Reload()
}
