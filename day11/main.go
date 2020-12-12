package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"reflect"
	"strings"
)

type Direction struct {
	x, y int
}

var (
	Up            = Direction{x: -1, y: 0}
	UpRight       = Direction{x: -1, y: 1}
	Right         = Direction{x: 0, y: 1}
	DownRight     = Direction{x: 1, y: 1}
	Down          = Direction{x: 1, y: 0}
	DownLeft      = Direction{x: 1, y: -1}
	Left          = Direction{x: 0, y: -1}
	UpLeft        = Direction{x: -1, y: -1}
	AllDirections = []Direction{Up, UpRight, Right, DownRight, Down, DownLeft, Left, UpLeft}
)

func (d *Direction) Sum(x, y int) (int, int) {
	x += d.x
	y += d.y
	return x, y
}

type SeatsMap struct {
	Map    [][]string
	width  int
	height int
}

func (s *SeatsMap) Width() int {
	return s.width
}

func (s *SeatsMap) Height() int {
	return s.height
}

func (s *SeatsMap) CountOccupiedNeighbors(x, y int) int {
	count := 0
	for _, d := range AllDirections {
		xi, yi := d.Sum(x, y)
		if xi < 0 || xi > s.height-1 {
			continue
		}
		if yi < 0 || yi > s.width-1 {
			continue
		}
		v := s.Map[xi][yi]
		switch v {
		case "L":
			continue
		case ".":
			continue
		case "#":
			count++
		}
	}
	return count
}

func (s *SeatsMap) CountOccupiedDirection(x, y int) int {
	count := 0
	for _, d := range AllDirections {
		xi := x
		yi := y
		for {
			xi, yi = d.Sum(xi, yi)
			if xi < 0 || xi > s.height-1 {
				goto next
			}
			if yi < 0 || yi > s.width-1 {
				goto next
			}
			v := s.Map[xi][yi]
			switch v {
			case "L":
				goto next
			case ".":
				continue
			case "#":
				count++
				goto next
			}
		}
	next:
	}
	return count
}

func (sm *SeatsMap) CountOccupied() int {
	count := 0
	for _, row := range sm.Map {
		for _, v := range row {
			if v == "#" {
				count++
			}
		}
	}
	return count
}

func (sm *SeatsMap) Print() string {
	buf := ""
	for _, row := range sm.Map {
		buf = fmt.Sprintf("%s%s\n", buf, strings.Join(row, ""))
	}
	return buf
}

func Duplicate(a [][]string) [][]string {
	n := len(a)
	m := len(a[0])
	duplicate := make([][]string, n)
	data := make([]string, n*m)
	for i := range a {
		start := i * m
		end := start + m
		duplicate[i] = data[start:end:end]
		copy(duplicate[i], a[i])
	}
	return duplicate
}

func RunIteration(sm *SeatsMap) bool {
	lm := Duplicate(sm.Map)
	for j, row := range sm.Map {
		for i, v := range row {
			if v == "." {
				continue
			}
			occupied := sm.CountOccupiedNeighbors(j, i)
			if v == "L" && occupied == 0 {
				lm[j][i] = "#"
			} else if v == "#" && occupied >= 4 {
				lm[j][i] = "L"
			}
		}
	}
	changed := reflect.DeepEqual(lm, sm.Map)
	sm.Map = lm
	fmt.Println(sm.Print())
	return !changed
}

func RunIterationV2(sm *SeatsMap) bool {
	lm := Duplicate(sm.Map)
	for j, row := range sm.Map {
		for i, v := range row {
			if v == "." {
				continue
			}
			occupied := sm.CountOccupiedDirection(j, i)
			if v == "L" && occupied == 0 {
				lm[j][i] = "#"
			} else if v == "#" && occupied >= 5 {
				lm[j][i] = "L"
			}
		}
	}
	changed := reflect.DeepEqual(lm, sm.Map)
	sm.Map = lm
	fmt.Println(sm.Print())
	return !changed
}

func RunUntilStable(sm *SeatsMap) int {
	count := 0
	for {
		if RunIteration(sm) {
			count++
		} else {
			break
		}
	}
	return count
}

func RunUntilStableV2(sm *SeatsMap) int {
	count := 0
	for {
		if RunIterationV2(sm) {
			count++
		} else {
			break
		}
	}
	return count
}

func LoadSeatMap(input io.ReadCloser) *SeatsMap {
	m := &SeatsMap{}
	s := bufio.NewScanner(input)
	defer input.Close()
	for s.Scan() {
		txt := strings.TrimSpace(s.Text())
		line := []string{}
		for _, v := range txt {
			line = append(line, string(v))
		}
		m.Map = append(m.Map, line)
	}
	m.width = len(m.Map[0])
	m.height = len(m.Map)
	fmt.Printf("w: %d h: %d\n", m.width, m.height)
	return m
}

func main() {
	filename := os.Args[1]
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	seats := LoadSeatMap(f)
	RunUntilStable(seats)

	fmt.Printf("Diffs: %#v\n", seats.CountOccupied())

	f, err = os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	seats = LoadSeatMap(f)
	RunUntilStableV2(seats)

	fmt.Printf("Diffs: %#v\n", seats.CountOccupied())
}
