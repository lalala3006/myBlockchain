package cli

import (
	"myBlockchain/block"
	"myBlockchain/network"
)

func (cli *Cli) genesis(address string, value int) {
	bc := block.NewBlockchain()
	bc.CreataGenesisTransaction(address, value, network.Send{})
}
