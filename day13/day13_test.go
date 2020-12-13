package main

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanLoadPuzzle(t *testing.T) {
	input := ioutil.NopCloser(bytes.NewReader([]byte(`939
7,13,x,x,59,x,31,19`)))

	puzzle := LoadPuzzle(input)

	assert.Equal(t, 939, puzzle.Timestamp)
	assert.Equal(t, []int{7, 13, 59, 31, 19}, puzzle.Buses)

	pairs := puzzle.ProcessBuses()
	assert.Equal(t, []Pair{
		{Bus: 7, Timestamp: 938},
		{Bus: 13, Timestamp: 936},
		{Bus: 59, Timestamp: 944},
		{Bus: 31, Timestamp: 930},
		{Bus: 19, Timestamp: 931},
	}, pairs)

	p := FindBusPair(puzzle.Timestamp, pairs)
	assert.Equal(t, Pair{Bus: 59, Timestamp: 944}, p)
}
