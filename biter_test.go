package biter

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func Test_zero(t *testing.T) {
	should := require.New(t)
	bits := Bits(0)
	iter := bits.ScanForward()
	should.Equal(64, iter())
	should.Equal(64, iter())
	iter = bits.ScanBackward()
	should.Equal(64, iter())
	should.Equal(64, iter())
}

func Test_one(t *testing.T) {
	should := require.New(t)
	bits := Bits(1)
	iter := bits.ScanForward()
	should.Equal(63, iter())
	should.Equal(64, iter())
	should.Equal(64, iter())
	iter = bits.ScanBackward()
	should.Equal(0, iter())
	should.Equal(64, iter())
}

func Test_two(t *testing.T) {
	should := require.New(t)
	bits := Bits(2)
	iter := bits.ScanForward() // 10
	should.Equal(62, iter())
	should.Equal(64, iter())
	iter = bits.ScanBackward()
	should.Equal(1, iter())
	should.Equal(64, iter())
}

func Test_three(t *testing.T) {
	should := require.New(t)
	bits := Bits(3)
	iter := bits.ScanForward() // 11
	should.Equal(62, iter())
	should.Equal(63, iter())
	should.Equal(64, iter())
	iter = bits.ScanBackward()
	should.Equal(0, iter())
	should.Equal(1, iter())
	should.Equal(64, iter())
}

func Test_9223372036854775809(t *testing.T) {
	should := require.New(t)
	bits := SetBits[0] | SetBits[63]
	iter := bits.ScanForward() // 1000000000000000000000000000000000000000000000000000000000000001
	should.Equal(0, iter())
	should.Equal(63, iter())
	should.Equal(64, iter())
	iter = bits.ScanBackward()
	should.Equal(0, iter())
	should.Equal(63, iter())
	should.Equal(64, iter())
}
