package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	TileIDMatcher  = regexp.MustCompile("Tile (?P<tileID>[0-9]*):")
	ClockWiseEdges = []int{0, 1, 2, 3}
)

type Neighbor struct {
	SelfEdge  int
	SelfFlip  bool
	OtherEdge int
	Tile      *Tile
}

type Tile struct {
	ID    int
	Data  [][]string
	Edges [][]string

	Neighbors []*Neighbor
}

type Image struct {
	Tiles   []*Tile
	Corners []*Tile
	Data    [][]string

	visited []*Tile
}

func NewImage() *Image {
	return &Image{
		Tiles:   []*Tile{},
		Corners: []*Tile{},
	}
}

func FlipSlice(input []string) []string {
	s := make([]string, len(input))
	for i, j := len(s)-1, 0; i >= 0; i-- {
		s[j] = input[i]
		j++
	}
	fmt.Printf("input: %v, s:%v\n", input, s)
	return s
}

func Equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func (im *Image) FindTile(ID int) *Tile {
	for _, t := range im.Tiles {
		if t.ID == ID {
			return t
		}
	}
	return nil
}

func (im *Image) FindCornerTiles() {
	for _, t := range im.Tiles {
		hasMatchs := 0
		for i, edge := range t.Edges {
			if ok, tn, ti := im.EdgeHasMatch(t.ID, edge); ok {
				//fmt.Printf("t.ID: %d e: %v matchs %d\n", t.ID, edge, ID)
				t.Neighbors = append(t.Neighbors, &Neighbor{SelfEdge: i, SelfFlip: false, OtherEdge: ti, Tile: tn})
				hasMatchs++
			} else if ok, tn, ti := im.EdgeHasMatch(t.ID, FlipSlice(edge)); ok {
				//fmt.Printf("t.ID: %d e fliped: %v matchs %d\n", t.ID, FlipSlice(edge), ID)
				t.Neighbors = append(t.Neighbors, &Neighbor{SelfEdge: i, SelfFlip: true, OtherEdge: ti, Tile: tn})
				hasMatchs++
			}
		}
		fmt.Printf("tileID: %d hasMatchs: %d\n", t.ID, hasMatchs)
		if hasMatchs == 2 {
			im.Corners = append(im.Corners, t)
		}
	}
}

func (im *Image) EdgeHasMatch(ID int, edge []string) (bool, *Tile, int) {
	for _, t := range im.Tiles {
		if t.ID == ID {
			continue
		}
		// fmt.Printf("ID: %d\n", t.ID)
		for i, e := range t.Edges {
			// fmt.Printf("edge: %#v => matchs: %#v?\n", edge, e)
			if Equal(edge, e) {
				return true, t, i
			}
		}
	}
	return false, nil, -1
}

func (im *Image) MultiplyCorners() int {
	r := 1
	for _, v := range im.Corners {
		r *= v.ID
	}
	return r
}

func TileInList(t *Tile, list []*Tile) bool {
	for _, tt := range list {
		if t.ID == tt.ID {
			return true
		}
	}
	return false
}

func (im *Image) PrintData(t0 *Tile, n *Neighbor) {
	if TileInList(n.Tile, im.visited) {
		return
	}
	fmt.Printf("%d [%d=>%d %v] => %d\n", t0.ID, n.SelfEdge, n.OtherEdge, n.SelfFlip, n.Tile.ID)
	im.visited = append(im.visited, n.Tile)
	for _, tt := range n.Tile.Neighbors {
		im.PrintData(n.Tile, tt)
	}
}

func (im *Image) Print() {
	t0 := im.Corners[0]
	fmt.Printf("T0: %d\n", t0.ID)
	im.visited = []*Tile{t0}
	for _, t := range t0.Neighbors {
		im.PrintData(t0, t)
	}
}

func (im *Image) LoadTiles(input io.ReadCloser) {
	s := bufio.NewScanner(input)
	defer input.Close()
	line := 0
	t := &Tile{}
	rEdge := []string{}
	lEdge := []string{}
	for s.Scan() {
		txt := strings.TrimSpace(s.Text())
		if txt == "" {
			im.Tiles = append(im.Tiles, t)
			t = &Tile{}
			rEdge = []string{}
			lEdge = []string{}
			continue
		}
		if strings.Contains(txt, "Tile") {
			parts := TileIDMatcher.FindStringSubmatch(txt)
			id, _ := strconv.Atoi(parts[1])
			t.ID = id
			line = 0
		} else {
			l := []string{}
			for _, v := range txt {
				l = append(l, string(v))
			}
			lEdge = append(lEdge, l[0])
			rEdge = append(rEdge, l[len(l)-1])

			t.Data = append(t.Data, l)
			if line == 0 {
				t.Edges = append(t.Edges, l)
			} else if line == 9 {
				t.Edges = append(t.Edges, rEdge, l, lEdge)
			}
			line++
		}
	}
	im.Tiles = append(im.Tiles, t)

}

func main() {
	filename := os.Args[1]
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	im := NewImage()
	im.LoadTiles(f)
	im.FindCornerTiles()
	im.Print()
	fmt.Printf("Result: %v\n", im.MultiplyCorners())
}
