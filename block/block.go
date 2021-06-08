package block

import (
	"bytes"
	"encoding/gob"
	log "github.com/corgi-kx/logcustom"
	"myBlockchain/database"
	"time"
)

type Block struct {
	//上一个区块的hash
	PreHash []byte
	//数据data
	Transactions []Transaction
	//时间戳
	TimeStamp int64
	//区块高度
	Height int
	//随机数
	Nonce int64
	//本区块hash
	Hash []byte
}

//进行挖矿来生成区块
func mineBlock(transaction []Transaction, preHash []byte, height int) (*Block, error) {
	timeStamp := time.Now().Unix()
	//hash数据+时间戳+上一个区块hash
	block := Block{preHash, transaction, timeStamp, height, 0, nil}
	pow := NewProofOfWork(&block)
	nonce, hash, err := pow.run()
	if err != nil {
		return nil, err
	}
	block.Nonce = nonce
	block.Hash = hash[:]
	log.Info("pow verify : ", pow.Verify())
	log.Infof("已生成新的区块,区块高度为%d", block.Height)
	return &block, nil
}

//添加区块信息到数据库，并更新lastHash
func (bc *blockchain) AddBlock(block *Block) {
	bc.BD.Put(block.Hash, block.Serialize(), database.BlockBucket)
	bci := NewBlockchainIterator(bc)
	currentBlock := bci.Next()
	if currentBlock == nil || currentBlock.Height < block.Height {
		bc.BD.Put([]byte(LastBlockHashMapping), block.Hash, database.BlockBucket)
	}
}

// 将Block对象序列化成[]byte
func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(b)
	if err != nil {
		panic(err)
	}
	return result.Bytes()
}

func (v *Block) Deserialize(d []byte) {
	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(v)
	if err != nil {
		log.Panic(err)
	}
}
