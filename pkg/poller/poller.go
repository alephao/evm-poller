package poller

import (
	"context"
	"log"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
)

func Poll(
	ctx context.Context,
	clients []*ethclient.Client,
	interval time.Duration,
	handler func(fromBlockNumber, toBlockNumber uint64, unpause func()),
) {
	ticker := NewTickerChannerWrapper(time.NewTicker(interval * time.Second))

	blockNumberGetters := make([]BlockNumberGetter, len(clients))
	for i, client := range clients {
		blockNumberGetters[i] = client
	}

	poll(ctx, blockNumberGetters, ticker, handler)
}

func poll(
	ctx context.Context,
	blockNumberGetters []BlockNumberGetter,
	ticker TickerChannelGetter,
	handler func(fromBlockNumber, toBlockNumber uint64, unpause func()),
) {
	var lastBlockNumber uint64
	for lastBlockNumber == 0 {
		lastBlockNumber = tryGettingBlockNumber(ctx, blockNumberGetters)
		if lastBlockNumber == 0 {
			time.Sleep(5 * time.Second)
		}
	}
	paused := true
	handler(lastBlockNumber, lastBlockNumber, func() {
		paused = false
	})
	for range ticker.TickerChannel() {
		if !paused {
			toBlockNumber := tryGettingBlockNumber(ctx, blockNumberGetters)

			// If no block number found, wait until next pull
			if toBlockNumber == 0 {
				continue
			}

			// If latest block is the same, wait until next pull
			if lastBlockNumber == toBlockNumber {
				continue
			}

			fromBlockNumber := lastBlockNumber + 1
			lastBlockNumber = toBlockNumber

			// pause so if another pull request comes, it is ignored until
			// the work is unpaused by whoever is using this functions
			paused = true
			handler(fromBlockNumber, toBlockNumber, func() {
				paused = false
			})
		}
	}
}

// Try getting the latest block number on each client, if not successful, returns 0
func tryGettingBlockNumber(ctx context.Context, blockNumberGetters []BlockNumberGetter) uint64 {
	for _, blockNumberGetter := range blockNumberGetters {
		toBlockNumber, err := blockNumberGetter.BlockNumber(ctx)
		if err != nil {
			log.Printf("failed to get block: %s", err.Error())
			continue
		}
		return toBlockNumber
	}
	return 0
}
