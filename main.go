package main

import (
	"context"
	"log"
	"time"

	"github.com/alephao/evm-poller/pkg/poller"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	rpcUrls := []string{
		"https://eth-mainnet.g.alchemy.com/v2/<key>",
		"https://mainnet.infura.io/v3/<key>",
		"https://<name>.discover.quiknode.pro/<key>/",
	}
	interval := 20 // polling interval in seconds

	log.Println("Connecting to ethereum clients...")
	ethClients := []*ethclient.Client{}
	for _, endpoint := range rpcUrls {
		client, err := ethclient.Dial(endpoint)
		if err != nil {
			log.Fatalf("failed to dial to RPC: %s\n", err.Error())
		}
		ethClients = append(ethClients, client)
	}

	poller.Poll(context.Background(), ethClients, time.Duration(interval), func(fromBlockNumber, toBlockNumber uint64, unpause func()) {
		log.Printf("Block Numbers: %d to %d", fromBlockNumber, toBlockNumber)
		unpause() // Don't forget this
	})
}
