package cli

import (
	"bufio"
	"fmt"
	log "github.com/corgi-kx/logcustom"
	"myBlockchain/network"
	"os"
	"strings"
	"time"
)

type Cli struct {
}

func New() *Cli {
	return &Cli{}
}

func (cli *Cli) Run() {
	printUsage()
	go cli.startNode()
	cli.ReceiveCMD()
}

//打印帮助提示
func printUsage() {
	fmt.Println("----------------------------------------------------------------------------- ")
	fmt.Println("Usage:")
	fmt.Println("\thelp                                              打印命令行说明")
	fmt.Println("\tquit                                              退出网络")
	fmt.Println("\ttest                                              测试")
	fmt.Println("------------------------------------------------------------------------------")
}

//获取用户输入
func (cli Cli) ReceiveCMD() {
	stdReader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		sendData, err := stdReader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading from stdin")
			panic(err)
		}
		cli.userCmdHandle(sendData)
	}
}

//用户输入命令的解析
func (cli Cli) userCmdHandle(data string) {
	//去除命令前后空格
	data = strings.TrimSpace(data)
	switch data {
	case "help":
		printUsage()
	case "quit":
		network.Send{}.SendSignOutToPeers()
		fmt.Println("本地节点已退出")
		time.Sleep(time.Second)
		os.Exit(0)
	case "test":
		log.Info("测试向log文件中添加信息")
	default:
		fmt.Println("无此命令!")
		printUsage()
	}
}
