package network

import (
	log "github.com/corgi-kx/logcustom"
	"github.com/libp2p/go-libp2p-core/network"
	"io/ioutil"
)

//对接收到的数据解析出命令,然后对不同的命令分别进行处理
func handleStream(stream network.Stream) {
	_, err := ioutil.ReadAll(stream)
	if err != nil {
		log.Panic(err)
	}

	log.Info("handle 函数")
}
