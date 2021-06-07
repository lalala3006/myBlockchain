package block

//当前本地监听端口
var ListenPort string

//钱包地址在数据库中的键
const addrListMapping = "addressList"

//中文助记词地址
var ChineseMnwordPath string

//公链版本信息默认为0
const version = byte(0x00)

//两次sha256(公钥hash)后截取的字节数量
const checkSum = 4
