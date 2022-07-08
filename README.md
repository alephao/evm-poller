# Eth Poller

Poll a range of block numbers from an evm node every x seconds

### Example

```go
package main

import (
	"context"
	"log"
	"time"

	"github.com/alephao/evm-poller/pkg/poller"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	rpcUrl := "https://eth-mainnet.g.alchemy.com/v2/<key>"
	interval := 20 // polling interval in seconds

	log.Println("Connecting to ethereum...")
	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		log.Fatalf("failed to dial to RPC: %s", err.Error())
	}

	poller.Poll(context.Background(), client, time.Duration(interval), func(fromBlockNumber, toBlockNumber uint64) {
		log.Printf("Block Numbers: %d to %d", fromBlockNumber, toBlockNumber)
	})
}
```