package block

import "myBlockchain/database"

type blockchain struct {
	BD *database.BlockchainDB //封装的blot结构体
}

//创建区块链实例
func NewBlockchain() *blockchain {
	blockchain := blockchain{}
	bd := database.New()
	blockchain.BD = bd
	return &blockchain
}
