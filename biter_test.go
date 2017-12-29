package biter

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func Test_zero(t *testing.T) {
	should := require.New(t)
	iter := ScanForward(0)
	should.Equal(64, iter())
	should.Equal(64, iter())
	iter = ScanBackward(0)
	should.Equal(64, iter())
	should.Equal(64, iter())
}

func Test_one(t *testing.T) {
	should := require.New(t)
	iter := ScanForward(1)
	should.Equal(63, iter())
	should.Equal(64, iter())
	should.Equal(64, iter())
	iter = ScanBackward(1)
	should.Equal(0, iter())
	should.Equal(64, iter())
}

func Test_two(t *testing.T) {
	should := require.New(t)
	iter := ScanForward(2) // 10
	should.Equal(62, iter())
	should.Equal(64, iter())
	iter = ScanBackward(2)
	should.Equal(1, iter())
	should.Equal(64, iter())
}

func Test_three(t *testing.T) {
	should := require.New(t)
	iter := ScanForward(3) // 11
	should.Equal(62, iter())
	should.Equal(63, iter())
	should.Equal(64, iter())
	iter = ScanBackward(3)
	should.Equal(0, iter())
	should.Equal(1, iter())
	should.Equal(64, iter())
}

func Test_9223372036854775809(t *testing.T) {
	should := require.New(t)
	iter := ScanForward(9223372036854775809) // 1000000000000000000000000000000000000000000000000000000000000001
	should.Equal(0, iter())
	should.Equal(63, iter())
	should.Equal(64, iter())
	iter = ScanBackward(9223372036854775809)
	should.Equal(0, iter())
	should.Equal(63, iter())
	should.Equal(64, iter())
}