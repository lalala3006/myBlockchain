package block

import (
	"bytes"
	"encoding/gob"
	log "github.com/corgi-kx/logcustom"
	"myBlockchain/database"
)

//UTXO输出的详细信息,便于直接在utxo数据库查找输出
type UTXO struct {
	Hash  []byte
	Index int
	Vout  TXOutput
}

type UTXOHandle struct {
	BC *blockchain
}

//重置UTXO数据库
func (u *UTXOHandle) ResetUTXODataBase() {
	//先查找全部未花费UTXO
	utxosMap := u.BC.findAllUTXOs()
	if utxosMap == nil {
		log.Debug("找不到区块,暂不重置UTXO数据库")
		return
	}
	//删除旧的UTXO数据库
	if database.IsBucketExist(u.BC.BD, database.UTXOBucket) {
		u.BC.BD.DeleteBucket(database.UTXOBucket)
	}
	//创建并将未花费UTXO循环添加
	for k, v := range utxosMap {
		u.BC.BD.Put([]byte(k), u.serialize(v), database.UTXOBucket)
	}
}

//查找数据库中全部未花费的UTXO
func (bc *blockchain) findAllUTXOs() map[string][]*UTXO {
	utxosMap := make(map[string][]*UTXO)
	txInputmap := make(map[string][]TXInput)
	bcIterator := NewBlockchainIterator(bc)
	for {
		currentBlock := bcIterator.Next()
		if currentBlock == nil {
			return nil
		}
		//必须倒序 否则有的已花费不会被扣掉
		for i := len(currentBlock.Transactions) - 1; i >= 0; i-- {
			var utxos = []*UTXO{}
			ts := currentBlock.Transactions[i]
			for _, vInt := range ts.Vint {
				txInputmap[string(vInt.TxHash)] = append(txInputmap[string(vInt.TxHash)], vInt)
			}

		VoutTag:
			for index, vOut := range ts.Vout {
				if txInputmap[string(ts.TxHash)] == nil {
					utxos = append(utxos, &UTXO{ts.TxHash, index, vOut})
				} else {
					for _, vIn := range txInputmap[string(ts.TxHash)] {
						if vIn.Index == index {
							continue VoutTag
						}
					}
					utxos = append(utxos, &UTXO{ts.TxHash, index, vOut})
				}
				utxosMap[string(ts.TxHash)] = utxos
			}
		}

		if isGenesisBlock(currentBlock) {
			break
		}
	}
	return utxosMap
}

func (u *UTXOHandle) serialize(utxos []*UTXO) []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(utxos)
	if err != nil {
		panic(err)
	}
	return result.Bytes()
}

func (u *UTXOHandle) dserialize(d []byte) []*UTXO {
	var model []*UTXO
	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&model)
	if err != nil {
		log.Panic(err)
	}
	return model
}
