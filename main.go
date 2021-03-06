package main

import (
	"fmt"
	log "github.com/corgi-kx/logcustom"
	"github.com/spf13/viper"
	"myBlockchain/block"
	"myBlockchain/cli"
	"myBlockchain/database"
	"myBlockchain/network"
	"os"
)

func init() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	logPath := viper.GetString("constant.log_path")
	listenHost := viper.GetString("network.listen_host")
	listenPort := viper.GetString("network.listen_port")
	rendezvousString := viper.GetString("network.rendezvous_string")
	protocolID := viper.GetString("network.protocol_id")
	mineDifficultyValue := viper.GetInt("blockchain.mine_difficulty_value")
	chineseMnwordPath := viper.GetString("blockchain.chinese_mnemonic_path")

	network.ListenHost = listenHost
	network.RendezvousString = rendezvousString
	network.ProtocolID = protocolID
	network.ListenPort = listenPort

	database.ListenPort = listenPort
	block.ListenPort = listenPort
	block.TargetBits = uint(mineDifficultyValue)
	block.ChineseMnwordPath = chineseMnwordPath

	//将日志输出到指定文件
	file, err := os.OpenFile(fmt.Sprintf("%slog%s.txt", logPath, listenPort), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Error(err)
	}
	log.SetOutputAll(file)

}

func main() {
	c := cli.New()
	c.Run()
}
