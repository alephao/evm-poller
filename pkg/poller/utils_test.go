package poller

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinUint64(t *testing.T) {
	testCases := [][]uint64{
		{10, 20, 10},
		{20, 10, 10},
		{20, 20, 20},
	}

	for _, testCase := range testCases {
		a := testCase[0]
		b := testCase[1]
		expected := testCase[2]
		assert.Equal(t, minUint64(a, b), expected)
	}
}
