package database

var ListenPort string

// 仓库类型
type BucketType string

const (
	BlockBucket BucketType = "blocks"
	AddrBucket  BucketType = "address"
	UTXOBucket  BucketType = "utxo"
)
