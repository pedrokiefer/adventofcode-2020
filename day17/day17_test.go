package main

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanLoadInitialMap(t *testing.T) {
	input := ioutil.NopCloser(bytes.NewReader([]byte(`.#.
..#
###`)))

	c := NewCube()
	c.LoadInitialMap(input)

	assert.True(t, c.Positions[Position{X: 0, Y: 1, Z: 0}])
}

func TestCanRunCycle(t *testing.T) {
	input := ioutil.NopCloser(bytes.NewReader([]byte(`.#.
..#
###`)))

	c := NewCube()
	c.LoadInitialMap(input)

	for i := 0; i < 6; i++ {
		c.RunCycle()
		c.Print()
	}

	assert.Equal(t, 848, len(c.Positions))
}
