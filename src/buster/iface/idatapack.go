package iface

/**
封包 拆包模块
*/

type IDataPack interface {
	//获取包头长度
	GetHeadLen() uint32

	//封包
	Pack(msg IMessage) (data []byte, err error)

	//拆包
	UnPack(data []byte) (msg IMessage, err error)
}
