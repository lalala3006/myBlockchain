package block

import "math"

//当前本地监听端口
var ListenPort string

//中文助记词地址
var ChineseMnwordPath string

//公链版本信息默认为0
const version = byte(0x00)

//两次sha256(公钥hash)后截取的字节数量
const checkSum = 4

//钱包地址在数据库中的键
const addrListMapping = "addressList"

//最新区块Hash在数据库中的键
const LastBlockHashMapping = "lastHash"

//当前节点发现的网络中最新区块高度
var NewestBlockHeight int

//随机数不能超过的最大值
const maxInt = math.MaxInt64

//挖矿难度值
var TargetBits uint
