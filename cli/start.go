package cli

import (
	"bufio"
	"fmt"
	log "github.com/corgi-kx/logcustom"
	"myBlockchain/network"
	"os"
	"strconv"
	"strings"
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
	fmt.Println("\tgenerateWallet                                    创建新钱包")
	fmt.Println("\tgenesis  -a DATA  -v DATA                         生成创世区块")
	fmt.Println("\tquit                                              退出网络")
	fmt.Println("\ttest -t TYPE                                      测试")
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
	var cmd string
	//var context string
	if strings.Contains(data, " ") {
		cmd = data[:strings.Index(data, " ")]
		//context = data[strings.Index(data, " ")+1:]
	} else {
		cmd = data
	}
	switch cmd {
	case "help":
		printUsage()
	case "generateWallet":
		cli.generateWallet()
	case "genesis":
		address := getSpecifiedContent(data, "-a", "-v")
		value := getSpecifiedContent(data, "-v", "")
		v, err := strconv.Atoi(value)
		if err != nil {
			log.Fatal(err)
		}
		cli.genesis(address, v)
	case "quit":
		flag := network.Send{}.SendSignOutToPeers()
		fmt.Println("本地节点已退出")
		if flag {
			os.Exit(0)
		}
	case "test":
		_type := getSpecifiedContent(data, "-t", "")
		cli.testCmd(_type)
	default:
		fmt.Println("无此命令!")
		printUsage()
	}
}

//返回data字符串中,标签为tag的内容
func getSpecifiedContent(data, tag, end string) string {
	if end != "" {
		return strings.TrimSpace(data[strings.Index(data, tag)+len(tag) : strings.Index(data, end)])
	}
	return strings.TrimSpace(data[strings.Index(data, tag)+len(tag):])
}
