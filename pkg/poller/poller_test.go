package poller

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPoll(t *testing.T) {
	blockNumberGetterMock := NewBlockNumberGetterMock()
	tickerMock := NewTickerChannelMock()

	wait := make(chan bool, 1)
	assertions := 0
	var expectedFrom uint64 = 0
	var expectedTo uint64 = 0
	go func() {
		poll(context.Background(), blockNumberGetterMock, tickerMock, func(fromBlockNumber, toBlockNumber uint64) {
			assert.Equal(t, fromBlockNumber, expectedFrom)
			assert.Equal(t, toBlockNumber, expectedTo)
			assertions++
			wait <- true
		})
	}()

	testCases := [][]uint64{
		{10, 10, 10},
		{20, 11, 20},
		{21, 21, 21},
		{30, 22, 30},
	}

	for i, testCase := range testCases {
		blockNumberGetterMock.BlockNumberResult = testCase[0]
		expectedFrom = testCase[1]
		expectedTo = testCase[2]
		tickerMock.Tick()
		<-wait
		assert.Equal(t, i+1, assertions)
	}
}
