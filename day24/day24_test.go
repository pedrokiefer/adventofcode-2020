package main

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanParseMoves(t *testing.T) {
	m := GetMoves("esenee")
	assert.Equal(t, []Move{East, SouthEast, NorthEast, East}, m)
}

func TestCanListOfTiles(t *testing.T) {
	input := ioutil.NopCloser(bytes.NewReader([]byte(`sesenwnenenewseeswwswswwnenewsewsw
neeenesenwnwwswnenewnwwsewnenwseswesw
seswneswswsenwwnwse
nwnwneseeswswnenewneswwnewseswneseene
swweswneswnenwsewnwneneseenw
eesenwseswswnenwswnwnwsewwnwsene
sewnenenenesenwsewnenwwwse
wenwwweseeeweswwwnwwe
wsweesenenewnwwnwsenewsenwwsesesenwne
neeswseenwwswnwswswnw
nenwswwsewswnenenewsenwsenwnesesenew
enewnwewneswsewnwswenweswnenwsenwsw
sweneswneswneneenwnewenewwneswswnese
swwesenesewenwneswnwwneseswwne
enesenwswwswneneswsenwnewswseenwsese
wnwnesenesenenwwnenwsewesewsesesew
nenewswnwewswnenesenwnesewesw
eneswnwswnwsenenwnwnwwseeswneewsenese
neswnwewnwnwseenwseesewsenwsweewe
wseweeenwnesenwwwswnew`)))

	tiles := LoadTiles(input)

	f := NewFloor()
	f.WalkAllTiles(tiles)
	f.Bounds()

	assert.Equal(t, 20, len(tiles))
	assert.Equal(t, 10, len(f.Tiles))

	f.DayIteration()
	assert.Equal(t, 15, len(f.Tiles))

	// 	f.DayIteration()
	// 	assert.Equal(t, 12, len(f.Tiles))

	// 	f.DayIteration()
	// 	assert.Equal(t, 25, len(f.Tiles))
	// }
}
