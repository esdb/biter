package biter

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func Test_zero(t *testing.T) {
	should := require.New(t)
	bits := SetBits[0]
	iter := bits.ScanForward()
	should.Equal(Slot(0), iter())
	should.Equal(NotFound, iter())
	iter = bits.ScanBackward()
	should.Equal(Slot(63), iter())
	should.Equal(NotFound, iter())
}

func Test_one(t *testing.T) {
	should := require.New(t)
	bits := SetBits[1]
	iter := bits.ScanForward()
	should.Equal(Slot(1), iter())
	should.Equal(NotFound, iter())
	should.Equal(NotFound, iter())
	iter = bits.ScanBackward()
	should.Equal(Slot(62), iter())
	should.Equal(NotFound, iter())
}

func Test_two(t *testing.T) {
	should := require.New(t)
	bits := SetBits[2]
	iter := bits.ScanForward() // 10
	should.Equal(Slot(2), iter())
	should.Equal(NotFound, iter())
	iter = bits.ScanBackward()
	should.Equal(Slot(61), iter())
	should.Equal(NotFound, iter())
}

func Test_zero_one_two(t *testing.T) {
	should := require.New(t)
	bits := SetBits[0] | SetBits[1] | SetBits[2]
	iter := bits.ScanForward() // 11
	should.Equal(Slot(0), iter())
	should.Equal(Slot(1), iter())
	should.Equal(Slot(2), iter())
	should.Equal(NotFound, iter())
	iter = bits.ScanBackward()
	should.Equal(Slot(61), iter())
	should.Equal(Slot(62), iter())
	should.Equal(Slot(63), iter())
	should.Equal(NotFound, iter())
}

func Test_9223372036854775809(t *testing.T) {
	should := require.New(t)
	bits := SetBits[0] | SetBits[63]
	iter := bits.ScanForward() // 1000000000000000000000000000000000000000000000000000000000000001
	should.Equal(Slot(0), iter())
	should.Equal(Slot(63), iter())
	should.Equal(NotFound, iter())
	iter = bits.ScanBackward()
	should.Equal(Slot(0), iter())
	should.Equal(Slot(63), iter())
	should.Equal(NotFound, iter())
}
