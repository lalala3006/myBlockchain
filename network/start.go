package network

import (
	"context"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
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

var PeerPool = make(map[string]peer.AddrInfo)
var ctx = context.Background()
var send = Send{}

//发送数据的头部多少位为命令
const prefixCMDLength = 12

//版本信息 默认0
const versionInfo = byte(0x00)

//网络通讯互相发送的命令
type command string

const (
	cVersion command = "version"
	//cGetHash     command = "getHash"
	//cHashMap     command = "hashMap"
	//cGetBlock    command = "getBlock"
	//cBlock       command = "block"
	//cTransaction command = "transaction"
	cMyError command = "myError"
)
