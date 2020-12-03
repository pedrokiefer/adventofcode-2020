package main

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanLoadTreeMap(t *testing.T) {
	input := ioutil.NopCloser(bytes.NewReader([]byte(`..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`)))

	m := LoadTreeMap(input)

	assert.Equal(t, 11, m.Height())
}

func TestCanCountTrees(t *testing.T) {
	input := ioutil.NopCloser(bytes.NewReader([]byte(`..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`)))

	m := LoadTreeMap(input)
	c := CountTrees(m, &Slope{Right: 3, Down: 1})

	assert.Equal(t, 11, m.Height())
	assert.Equal(t, 7, c)
}

func TestCanCountTreesMultipleSlopes(t *testing.T) {
	input := ioutil.NopCloser(bytes.NewReader([]byte(`..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`)))

	slopes := []*Slope{
		{Right: 1, Down: 1},
		{Right: 3, Down: 1},
		{Right: 5, Down: 1},
		{Right: 7, Down: 1},
		{Right: 1, Down: 2},
	}

	m := LoadTreeMap(input)
	c := CountSlopes(m, slopes)

	assert.Equal(t, 11, m.Height())
	assert.Equal(t, 336, c)
}
