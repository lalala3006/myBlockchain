package network

import (
	"github.com/libp2p/go-libp2p-core/host"
)

//p2p相关,程序启动时,会被配置文件所替换
var (
	RendezvousString = "meetme"
	ProtocolID       = "/chain/1.1.0"
	ListenHost       = "0.0.0.0"
	ListenPort       = "3001"
	localHost        host.Host
	localAddr        string
)
