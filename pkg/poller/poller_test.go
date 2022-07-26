package poller

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPoll(t *testing.T) {
	blockNumberGetterFailingMock := NewBlockNumberGetterMock()
	blockNumberGetterFailingMock.BlockNumberError = fmt.Errorf("error")

	blockNumberGetterMock := NewBlockNumberGetterMock()
	blockNumberGetterMock.BlockNumberResult = 10
	tickerMock := NewTickerChannelMock()

	wait := make(chan bool, 1)
	assertions := 0
	var expectedFrom uint64 = 10
	var expectedTo uint64 = 10
	go func() {
		poll(
			context.Background(),
			[]BlockNumberGetter{blockNumberGetterFailingMock, blockNumberGetterMock},
			tickerMock,
			func(fromBlockNumber, toBlockNumber uint64, unpause func(),
			) {
				assert.Equal(t, fromBlockNumber, expectedFrom)
				assert.Equal(t, toBlockNumber, expectedTo)
				assertions++
				wait <- true
				unpause()
			})
	}()
	<-wait

	testCases := [][]uint64{
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
		assert.Equal(t, i+2, assertions)
	}
}
