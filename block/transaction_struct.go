package block

//UTXO输入
type TXInput struct {
	TxHash    []byte
	Index     int
	Signature []byte
	PublicKey []byte
}

//UTXO输出
type TXOutput struct {
	Value         int
	PublicKeyHash []byte
}

//交易列表信息
type Transaction struct {
	TxHash []byte
	//UTXO输入
	Vint []TXInput
	//UTXO输出
	Vout []TXOutput
}
