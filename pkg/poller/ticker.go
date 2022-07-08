package poller

import "time"

type TickerChannelGetter interface {
	TickerChannel() <-chan time.Time
}

type TicketChannelWrapper struct {
	Ticker *time.Ticker
}

func NewTickerChannerWrapper(ticker *time.Ticker) *TicketChannelWrapper {
	return &TicketChannelWrapper{
		ticker,
	}
}

func (tcw TicketChannelWrapper) TickerChannel() <-chan time.Time {
	return tcw.Ticker.C
}

type TickerChannelMock struct {
	C chan time.Time
}

func NewTickerChannelMock() *TickerChannelMock {
	return &TickerChannelMock{
		C: make(chan time.Time),
	}
}

func (tcw TickerChannelMock) TickerChannel() <-chan time.Time {
	return tcw.C
}

func (tcw TickerChannelMock) Tick() {
	tcw.C <- time.Now()
}
