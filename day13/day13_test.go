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
	assert.Equal(t, []*Bus{
		{ID: 7, N: 0},
		{ID: 13, N: 1},
		{ID: 59, N: 4},
		{ID: 31, N: 6},
		{ID: 19, N: 7},
	}, puzzle.Buses)

	pairs := puzzle.ProcessBuses()
	assert.Equal(t, []Pair{
		{Bus: &Bus{ID: 7, N: 0}, Timestamp: 938},
		{Bus: &Bus{ID: 13, N: 1}, Timestamp: 936},
		{Bus: &Bus{ID: 59, N: 4}, Timestamp: 944},
		{Bus: &Bus{ID: 31, N: 6}, Timestamp: 930},
		{Bus: &Bus{ID: 19, N: 7}, Timestamp: 931},
	}, pairs)

	p := FindBusPair(puzzle.Timestamp, pairs)
	assert.Equal(t, Pair{Bus: &Bus{ID: 59, N: 4}, Timestamp: 944}, p)
}

func TestCanFindCommonTimestamp(t *testing.T) {
	ts := FindCommonTimestamp([]*Bus{{ID: 7, N: 0}, {ID: 13, N: 1}})
	assert.Equal(t, 77, ts)

	ts = FindCommonTimestamp([]*Bus{{ID: 7, N: 0}, {ID: 13, N: 1}, {ID: 59, N: 4}})
	assert.Equal(t, 350, ts)

	ts = FindCommonTimestamp([]*Bus{{ID: 7, N: 0}, {ID: 31, N: 6}})
	assert.Equal(t, 56, ts)

	ts = FindCommonTimestamp([]*Bus{{ID: 7, N: 0}, {ID: 13, N: 1}, {ID: 59, N: 4}, {ID: 31, N: 6}, {ID: 19, N: 7}})
	assert.Equal(t, 1068781, ts)
}
