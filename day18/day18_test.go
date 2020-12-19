package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanProcessALine(t *testing.T) {
	// 	input := ioutil.NopCloser(bytes.NewReader([]byte(`mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
	// mem[8] = 11
	// mem[7] = 101
	// mem[8] = 0`)))

	r := ProcessLine("1 + 2 * 3 + 4 * 5 + 6")
	assert.Equal(t, 71, r)

	r = ProcessLine("1 + (2 * 3) + (4 * (5 + 6))")
	assert.Equal(t, 51, r)

	r = ProcessLine("2 * 3 + (4 * 5)")
	assert.Equal(t, 26, r)

	r = ProcessLine("5 + (8 * 3 + 9 + 3 * 4 * 3)")
	assert.Equal(t, 437, r)

	r = ProcessLine("5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))")
	assert.Equal(t, 12240, r)

	r = ProcessLine("((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2")
	assert.Equal(t, 13632, r)
}

func TestCanProcessALineWithPrecedence(t *testing.T) {
	// 	input := ioutil.NopCloser(bytes.NewReader([]byte(`mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
	// mem[8] = 11
	// mem[7] = 101
	// mem[8] = 0`)))

	r := ProcessLinePrecedence("1 + 2 * 3 + 4 * 5 + 6")
	assert.Equal(t, 231, r)

	r = ProcessLinePrecedence("1 + (2 * 3) + (4 * (5 + 6))")
	assert.Equal(t, 51, r)

	r = ProcessLinePrecedence("2 * 3 + (4 * 5)")
	assert.Equal(t, 46, r)

	r = ProcessLinePrecedence("5 + (8 * 3 + 9 + 3 * 4 * 3)")
	assert.Equal(t, 1445, r)

	r = ProcessLinePrecedence("5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))")
	assert.Equal(t, 669060, r)

	r = ProcessLinePrecedence("((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2")
	assert.Equal(t, 23340, r)
}
