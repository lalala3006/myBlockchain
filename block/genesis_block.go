package block

import (
	"fmt"
	log "github.com/corgi-kx/logcustom"
	"math/big"
	"myBlockchain/database"
)

//创建创世区块交易信息
func (bc *blockchain) CreataGenesisTransaction(address string, value int, send Sender) {
	//判断地址格式是否正确
	if !IsVaildBitcoinAddress(address) {
		log.Errorf("地址格式不正确:%s\n", address)
		return
	}

	//创世区块数据
	txi := TXInput{[]byte{}, -1, nil, nil}
	//本地一定要存创世区块地址的公私钥信息
	wallets := NewWallets(bc.BD)
	genesisKeys, ok := wallets.Wallets[address]
	if !ok {
		log.Fatal("没有找到地址对应的公私钥信息")
	}
	//通过地址获得rip160(sha256(publickey))
	publicKeyHash := generatePublicKeyHash(genesisKeys.PublicKey)
	txo := TXOutput{value, publicKeyHash}
	ts := Transaction{nil, []TXInput{txi}, []TXOutput{txo}}
	ts.hash()
	tss := []Transaction{ts}

	//开始生成区块链的第一个区块
	flag := bc.newGenesisBlockchain(tss)
	if flag {
		fmt.Println("已成功生成创世区块\n")
	} else {
		fmt.Println("创世区块已存在，不可重复生成\n")
	}

	//创世区块后,更新本地最新区块为1并,向全网节点发送当前区块链高度1
	NewestBlockHeight = 1
	send.SendVersionToPeers(1)

	//重置utxo数据库，将创世数据存入
	utxos := UTXOHandle{bc}
	utxos.ResetUTXODataBase()
}

//创建区块链
func (bc *blockchain) newGenesisBlockchain(transaction []Transaction) bool {
	//判断一下是否已生成创世区块
	if len(bc.BD.View([]byte(LastBlockHashMapping), database.BlockBucket)) != 0 {
		//log.Fatal("不可重复生成创世区块")
		log.Error("不可重复生成创世区块")
		return false
	}
	//生成创世区块
	genesisBlock := newGenesisBlock(transaction)
	//添加到数据库中
	bc.AddBlock(genesisBlock)
	return true
}

//生成创世区块
func newGenesisBlock(transaction []Transaction) *Block {
	//创世区块的上一个块hash默认设置成下面的样子
	preHash := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	//生成创世区块
	genesisBlock, err := mineBlock(transaction, preHash, 1)
	if err != nil {
		log.Error(err)
	}
	return genesisBlock
}

func isGenesisBlock(block *Block) bool {
	var hashInt big.Int
	hashInt.SetBytes(block.PreHash)
	if big.NewInt(0).Cmp(&hashInt) == 0 {
		return true
	}
	return false
}
