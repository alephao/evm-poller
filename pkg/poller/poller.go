package poller

import (
	"context"
	"log"
	"math"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
)

func Poll(ctx context.Context, client *ethclient.Client, interval time.Duration, handler func(fromBlockNumber, toBlockNumber uint64)) {
	ticker := NewTickerChannerWrapper(time.NewTicker(interval * time.Second))
	poll(ctx, client, ticker, handler)
}

func poll(ctx context.Context, blockNumberGetter BlockNumberGetter, ticker TickerChannelGetter, handler func(fromBlockNumber, toBlockNumber uint64)) {
	var lastBlockNumber uint64 = math.MaxUint64 - 1_000_000

	for range ticker.TickerChannel() {
		toBlockNumber, err := blockNumberGetter.BlockNumber(ctx)
		if err != nil {
			log.Printf("failed to get block: %s", err.Error())
		}

		if lastBlockNumber == toBlockNumber {
			continue
		}

		fromBlockNumber := minUint64(lastBlockNumber+1, toBlockNumber)
		lastBlockNumber = toBlockNumber
		handler(fromBlockNumber, toBlockNumber)
	}
}
