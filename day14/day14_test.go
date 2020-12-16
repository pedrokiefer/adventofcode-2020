package main

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanGenerateMaskFunction(t *testing.T) {
	mask := GenMaskFunction("XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X")
	assert.Equal(t, uint64(73), mask(uint64(11)))
	assert.Equal(t, uint64(101), mask(uint64(101)))
	assert.Equal(t, uint64(64), mask(uint64(0)))
}

func TestCanGenAddressesFunction(t *testing.T) {
	mask := GenAddressesFunction("000000000000000000000000000000X1001X")
	assert.Equal(t, []uint64{26, 27, 58, 59}, mask(uint64(42)))

	mask = GenAddressesFunction("00000000000000000000000000000000X0XX")
	assert.Equal(t, []uint64{16, 17, 18, 19, 24, 25, 26, 27}, mask(uint64(26)))
}

func TestCanLoadNavigationInstructions(t *testing.T) {
	input := ioutil.NopCloser(bytes.NewReader([]byte(`mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
mem[8] = 11
mem[7] = 101
mem[8] = 0`)))

	m := NewMem()
	m.ProcessMemory(input, false)

	assert.Equal(t, 2, len(m.Addresses))
	assert.Equal(t, uint64(165), m.Sum())
}

func TestCanLoadNavigationInstructionsPart2(t *testing.T) {
	input := ioutil.NopCloser(bytes.NewReader([]byte(`mask = 000000000000000000000000000000X1001X
mem[42] = 100
mask = 00000000000000000000000000000000X0XX
mem[26] = 1`)))

	m := NewMem()
	m.ProcessMemory(input, true)

	assert.Equal(t, 10, len(m.Addresses))
	assert.Equal(t, uint64(208), m.Sum())
}
