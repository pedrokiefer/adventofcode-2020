package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type Point struct {
	X, Y int
}

func (p *Point) Walk(p1 Point) {
	p.X += p1.X
	p.Y += p1.Y
}

func (p *Point) Sum(p1 Point) Point {
	return Point{
		X: p.X + p1.X,
		Y: p.Y + p1.Y,
	}
}

type Move struct {
	Name      string
	Direction Point
}

var (
	East      = Move{Name: "e", Direction: Point{X: 2, Y: 0}}
	SouthEast = Move{Name: "se", Direction: Point{X: 1, Y: -1}}
	SouthWest = Move{Name: "sw", Direction: Point{X: -1, Y: -1}}
	West      = Move{Name: "w", Direction: Point{X: -2, Y: 0}}
	NorthWest = Move{Name: "nw", Direction: Point{X: -1, Y: 1}}
	NorthEast = Move{Name: "ne", Direction: Point{X: 1, Y: 1}}
	AllMoves  = []Move{East, SouthEast, SouthWest, West, NorthWest, NorthEast}
)

type Tile struct {
	Moves []Move
}

func GetMoves(txt string) []Move {
	ms := []Move{}
	for i := 0; i < len(txt); i++ {
		switch txt[i] {
		case 'e':
			ms = append(ms, East)
		case 'w':
			ms = append(ms, West)
		case 's':
			if txt[i+1] == 'e' {
				ms = append(ms, SouthEast)
			} else if txt[i+1] == 'w' {
				ms = append(ms, SouthWest)
			}
			i++
		case 'n':
			if txt[i+1] == 'e' {
				ms = append(ms, NorthEast)
			} else if txt[i+1] == 'w' {
				ms = append(ms, NorthWest)
			}
			i++
		}
	}
	return ms
}

func LoadTiles(input io.ReadCloser) []Tile {
	s := bufio.NewScanner(input)
	defer input.Close()
	ts := []Tile{}
	for s.Scan() {
		txt := strings.TrimSpace(s.Text())
		m := GetMoves(txt)
		ts = append(ts, Tile{Moves: m})
	}
	return ts
}

type Floor struct {
	Tiles      map[Point]bool
	Xmin, Xmax int
	Ymin, Ymax int
}

func NewFloor() Floor {
	return Floor{
		Tiles: map[Point]bool{},
		Xmin:  int(^uint(0) >> 1),
		Ymin:  int(^uint(0) >> 1),
		Xmax:  0,
		Ymax:  0,
	}
}

func (f *Floor) UpdateXMinMax(p Point) {
	if p.X < f.Xmin {
		f.Xmin = p.X
	}
	if p.X > f.Xmax {
		f.Xmax = p.X
	}
}

func (f *Floor) UpdateYMinMax(p Point) {
	if p.Y < f.Ymin {
		f.Ymin = p.Y
	}
	if p.Y > f.Ymax {
		f.Ymax = p.Y
	}
}

func (f *Floor) WalkPath(ms []Move) {
	p := Point{X: 0, Y: 0}
	for _, m := range ms {
		p.Walk(m.Direction)
	}

	if _, ok := f.Tiles[p]; ok {
		delete(f.Tiles, p)
	} else {
		f.Tiles[p] = true
	}
}

func (f *Floor) Bounds() {
	for p := range f.Tiles {
		f.UpdateXMinMax(p)
		f.UpdateYMinMax(p)
	}
}

func (f *Floor) WalkAllTiles(ts []Tile) {
	for _, t := range ts {
		f.WalkPath(t.Moves)
	}
}

func (f *Floor) CountNeighbors(p Point) int {
	count := 0
	for _, d := range AllMoves {
		np := p.Sum(d.Direction)
		if _, ok := f.Tiles[np]; ok {
			count++
		}
	}
	return count
}

func ToEven(v int) int {
	if v%2 == 0 {
		return v
	}
	if v < 0 {
		return v - 1
	}
	return v + 1
}

func IsValidHexGrid(p Point) bool {
	x := ((p.X % 2) + 2) % 2
	y := ((p.Y % 2) + 2) % 2
	if x == 0 && y == 0 {
		return true
	} else if x == 1 && y == 1 {
		return true
	}
	return false
}

func (f *Floor) DayIteration() {
	ts := map[Point]bool{}

	Xmin := ToEven(f.Xmin - 4)
	Xmax := ToEven(f.Xmax + 4)
	Ymin := ToEven(f.Ymin - 4)
	Ymax := ToEven(f.Ymax + 4)

	for i := Xmin; i <= Xmax; i++ {
		for j := Ymin; j <= Ymax; j++ {
			p := Point{X: i, Y: j}
			if !IsValidHexGrid(p) {
				continue
			}
			_, black := f.Tiles[p]
			c := f.CountNeighbors(p)

			if black && (c > 0 && c <= 2) {
				ts[p] = true
			} else if !black && c == 2 {
				ts[p] = true
			}
		}
	}
	f.Tiles = ts
	f.Bounds()
}

func main() {
	filename := os.Args[1]
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	tiles := LoadTiles(f)
	floor := NewFloor()
	floor.WalkAllTiles(tiles)
	floor.Bounds()

	fmt.Printf("Result: %d\n", len(floor.Tiles))

	for i := 0; i < 100; i++ {
		floor.DayIteration()
	}
	fmt.Printf("Result: %d\n", len(floor.Tiles))
}
