package cli

import "myBlockchain/network"

func (cli Cli) startNode() {
	network.StartNode(cli)
}
