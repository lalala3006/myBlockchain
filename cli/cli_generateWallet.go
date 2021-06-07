package cli

import (
	"errors"
	"fmt"
	"github.com/boltdb/bolt"
	log "github.com/corgi-kx/logcustom"
	"myBlockchain/block"
	"myBlockchain/database"
)

func (cli *Cli) generateWallet() {
	bd := database.New()
	wallets := block.NewWallets(bd)
	address, privkey, mnemonicWord := wallets.GenerateWallet(bd, block.NewBitcoinKeys, []string{})
	fmt.Println("助记词：", mnemonicWord)
	fmt.Println("私钥：", privkey)
	fmt.Println("地址：", address)
}

// test命令显示数据库
func (cli *Cli) testCmd() {
	var DBFileName = "blockchain_9001.db"
	db, err := bolt.Open(DBFileName, 0600, nil)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	bucketName := database.AddrBucket
	err = db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		if bucket == nil {
			msg := "datebase view warnning:没有找到仓库：" + string(bucketName)
			log.Error(msg)
			return errors.New(msg)
		}
		bucket.ForEach(func(k, v []byte) error {
			//log.Infof("key=%s, value=%s\n", string(k), string(v))
			log.Infof("key=%s", string(k))
			return nil
		})
		return nil
	})
}
