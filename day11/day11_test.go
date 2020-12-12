package main

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanReadSeatMap(t *testing.T) {
	input := ioutil.NopCloser(bytes.NewReader([]byte(`L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL`)))

	seats := LoadSeatMap(input)

	assert.Equal(t, 10, seats.Width())
	assert.Equal(t, 10, seats.Height())

	assert.Equal(t, 0, seats.CountOccupiedNeighbors(0, 0))
	assert.Equal(t, 0, seats.CountOccupiedNeighbors(10, 10))
	assert.Equal(t, 0, seats.CountOccupiedNeighbors(10, 10))
}

func TestCanReadSeatMapOccupied(t *testing.T) {
	input := ioutil.NopCloser(bytes.NewReader([]byte(`#.LL.LL.LL
##LLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.#
L.LLLLL.#L`)))

	seats := LoadSeatMap(input)

	assert.Equal(t, 10, seats.Width())
	assert.Equal(t, 10, seats.Height())

	assert.Equal(t, 2, seats.CountOccupiedNeighbors(0, 0))
	assert.Equal(t, 2, seats.CountOccupiedNeighbors(9, 9))
}

func TestCanRunInteration(t *testing.T) {
	input := ioutil.NopCloser(bytes.NewReader([]byte(`L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL`)))

	seats := LoadSeatMap(input)

	assert.True(t, RunIteration(seats))

	assert.Equal(t, 3, seats.CountOccupiedNeighbors(1, 0))

	assert.True(t, RunIteration(seats))

	assert.Equal(t, `#.LL.L#.##
#LLLLLL.L#
L.L.L..L..
#LLL.LL.L#
#.LL.LL.LL
#.LLLL#.##
..L.L.....
#LLLLLLLL#
#.LLLLLL.L
#.#LLLL.##
`, seats.Print())
}

func TestCanRunUntilStable(t *testing.T) {
	input := ioutil.NopCloser(bytes.NewReader([]byte(`L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL`)))

	seats := LoadSeatMap(input)

	assert.Equal(t, 5, RunUntilStable(seats))
	assert.Equal(t, 37, seats.CountOccupied())
}

func TestCanReadSeatMapV2(t *testing.T) {
	input := ioutil.NopCloser(bytes.NewReader([]byte(`.......#.
...#.....
.#.......
.........
..#L....#
....#....
.........
#........
...#.....`)))

	seats := LoadSeatMap(input)

	assert.Equal(t, 9, seats.Width())
	assert.Equal(t, 9, seats.Height())

	assert.Equal(t, 8, seats.CountOccupiedDirection(4, 3))
}

func TestCanRunInterationV2(t *testing.T) {
	input := ioutil.NopCloser(bytes.NewReader([]byte(`L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL`)))

	seats := LoadSeatMap(input)

	assert.True(t, RunIterationV2(seats))

	assert.Equal(t, 4, seats.CountOccupiedDirection(1, 0))

	assert.True(t, RunIterationV2(seats))

	assert.Equal(t, 0, seats.CountOccupiedDirection(1, 2))

	assert.True(t, RunIterationV2(seats))

	assert.Equal(t, `#.L#.##.L#
#L#####.LL
L.#.#..#..
##L#.##.##
#.##.#L.##
#.#####.#L
..#.#.....
LLL####LL#
#.L#####.L
#.L####.L#
`, seats.Print())
}

func TestCanRunUntilStableV2(t *testing.T) {
	input := ioutil.NopCloser(bytes.NewReader([]byte(`L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL`)))

	seats := LoadSeatMap(input)

	assert.Equal(t, 6, RunUntilStableV2(seats))
	assert.Equal(t, 26, seats.CountOccupied())
}
