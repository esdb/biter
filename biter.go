package biter

import (
	"math/bits"
	"math"
)

type Slot int
const NotFound Slot = 64
const SetAllBits = Bits(math.MaxUint64)

var SetBits []Bits
var SetBitsForwardFrom []Bits
var SetBitsForwardUntil []Bits

func init() {
	SetBits = make([]Bits, 64)
	SetBitsForwardFrom = make([]Bits, 64)
	SetBitsForwardUntil = make([]Bits, 64)
	for i := uint(0); i < 64; i++ {
		SetBits[i] = 1 << i
		for j := uint(0); j < i; j++ {
			SetBitsForwardUntil[i] |= 1 << j
		}
		for j := i; j < 64; j++ {
			SetBitsForwardFrom[i] |= 1 << j
		}
	}
}

type Bits uint64

func (b Bits) And(anotherBits Bits) Bits {
	return b & anotherBits
}

func (b Bits) Or(anotherBits Bits) Bits {
	return b | anotherBits
}

// from left to right
func (b Bits) ScanForward() func() Slot {
	lastPos := Slot(-1)
	return func() Slot {
		if b == 0 {
			return 64
		}
		trailingZeros := 1 + bits.TrailingZeros64(uint64(b))
		lastPos = lastPos + Slot(trailingZeros)
		b = b >> uint(trailingZeros)
		return lastPos
	}
}

// from right to left
func (b Bits) ScanBackward() func() Slot {
	lastPos := Slot(-1)
	return func() Slot {
		if b == 0 {
			return 64
		}
		leadingZeros := 1 + bits.LeadingZeros64(uint64(b))
		lastPos = lastPos + Slot(leadingZeros)
		b = b << uint(leadingZeros)
		return lastPos
	}
}
