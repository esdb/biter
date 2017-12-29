package biter

import (
	"math/bits"
)

// from left to right
func ScanForward(b uint64) func() int {
	lastPos := -1
	return func() int {
		if b == 0 {
			return 64
		}
		leadingZeros := 1 + bits.LeadingZeros64(b)
		lastPos = lastPos + leadingZeros
		b = b << uint(leadingZeros)
		return lastPos
	}
}

// from right to left
func ScanBackward(b uint64) func() int {
	lastPos := -1
	return func() int {
		if b == 0 {
			return 64
		}
		trailingZeros := 1 + bits.TrailingZeros64(b)
		lastPos = lastPos + trailingZeros
		b = b >> uint(trailingZeros)
		return lastPos
	}
}
