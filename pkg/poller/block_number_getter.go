package poller

import "context"

type BlockNumberGetter interface {
	BlockNumber(ctx context.Context) (uint64, error)
}

type BlockNumberGetterMock struct {
	BlockNumberResult uint64
	BlockNumberError  error
}

func NewBlockNumberGetterMock() *BlockNumberGetterMock {
	return &BlockNumberGetterMock{
		BlockNumberResult: 0,
		BlockNumberError:  nil,
	}
}

func (mock BlockNumberGetterMock) BlockNumber(ctx context.Context) (uint64, error) {
	return mock.BlockNumberResult, mock.BlockNumberError
}
