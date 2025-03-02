package main

import "github.com/NethermindEth/juno/feedergatewaysync/feedergatewaysync"

const (
	feeder_gateway_url = "https://alpha-sepolia.starknet.io"
	start_block        = 0
	end_block          = 100
)

func main() {
	sm := feedergatewaysync.NewSyncManager(feeder_gateway_url, 0, 100)
	sm.SyncBlocks()
}
