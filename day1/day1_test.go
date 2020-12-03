package main

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFileToIntList(t *testing.T) {
	input := ioutil.NopCloser(bytes.NewReader([]byte(`
1721
979
366
299
675
1456
`)))

	list := InputToIntList(input)

	assert.Equal(t, []int64{1721, 979, 366, 299, 675, 1456}, list)
}

func TestCanSum2To2020(t *testing.T) {
	input := []int64{1721, 979, 366, 299, 675, 1456}
	a, b := Sum2To2020(input)

	assert.Equal(t, int64(1721), a)
	assert.Equal(t, int64(299), b)
}

func TestCanSum3To2020(t *testing.T) {
	input := []int64{1721, 979, 366, 299, 675, 1456}
	a, b, c := Sum3To2020(input)

	assert.Equal(t, int64(979), a)
	assert.Equal(t, int64(366), b)
	assert.Equal(t, int64(675), c)
}
