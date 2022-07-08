package poller

func minUint64(a uint64, b uint64) uint64 {
	if a < b {
		return a
	}
	return b
}
