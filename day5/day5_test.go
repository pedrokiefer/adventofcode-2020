package main

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanGetSeatID(t *testing.T) {
	assert.Equal(t, 567, ToSeatID("BFFFBBFRRR"))
	assert.Equal(t, 119, ToSeatID("FFFBBBFRRR"))
	assert.Equal(t, 820, ToSeatID("BBFFBBFRLL"))
}

func TestCanFindHighestSeatID(t *testing.T) {
	input := ioutil.NopCloser(bytes.NewReader([]byte(`BFFFBBFRRR
FFFBBBFRRR
BBFFBBFRLL`)))

	IDs := LoadSeatIDs(input)

	assert.Equal(t, 820, IDs[len(IDs)-1])
}
