package network

import (
	"fmt"
	log "github.com/corgi-kx/logcustom"
	"github.com/libp2p/go-libp2p-core/network"
	"io/ioutil"
)

//对接收到的数据解析出命令,然后对不同的命令分别进行处理
func handleStream(stream network.Stream) {
	data, err := ioutil.ReadAll(stream)
	if err != nil {
		log.Panic(err)
	}

	//取信息的前十二位得到命令
	cmd, content := splitMessage(data)
	log.Tracef("本节点已接收到命令：%s", cmd)
	switch command(cmd) {
	//case cVersion:
	//	go handleVersion(content)
	//case cGetHash:
	//	go handleGetHash(content)
	//case cHashMap:
	//	go handleHashMap(content)
	//case cGetBlock:
	//	go handleGetBlock(content)
	//case cBlock:
	//	go handleBlock(content)
	//case cTransaction:
	//	go handleTransaction(content)
	case cMyError:
		go handleMyError(content)
	}
}

//打印接收到的错误信息
func handleMyError(content []byte) {
	e := myerror{}
	e.deserialize(content)
	log.Warn(e.Error)
	peer := buildPeerInfoByAddr(e.Addrfrom)
	delete(PeerPool, fmt.Sprint(peer.ID))
}
