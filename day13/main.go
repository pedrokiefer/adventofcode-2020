package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Bus struct {
	ID int
	N  int
}

type Pair struct {
	Bus       *Bus
	Timestamp int
}

type Puzzle struct {
	Timestamp int
	Buses     []*Bus
}

func LoadPuzzle(input io.ReadCloser) *Puzzle {
	puzzle := &Puzzle{
		Buses: []*Bus{},
	}
	defer input.Close()
	data, _ := ioutil.ReadAll(input)
	txt := string(data)
	lines := strings.Split(txt, "\n")
	ts, _ := strconv.Atoi(lines[0])
	puzzle.Timestamp = ts
	for i, v := range strings.Split(lines[1], ",") {
		if v == "x" {
			continue
		}
		bus, _ := strconv.Atoi(v)
		puzzle.Buses = append(puzzle.Buses, &Bus{ID: bus, N: i})
	}
	return puzzle
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (p *Puzzle) ProcessBuses() []Pair {
	result := []Pair{}
	for _, b := range p.Buses {
		v := float64(p.Timestamp) / float64(b.ID)
		f := int(math.Floor(v)) * b.ID
		c := int(math.Ceil(v)) * b.ID

		if Abs(c-p.Timestamp) < Abs(f-p.Timestamp) {
			result = append(result, Pair{Bus: b, Timestamp: c})
		} else {
			result = append(result, Pair{Bus: b, Timestamp: f})
		}
	}
	return result
}

func FindBusPair(ts int, pairs []Pair) Pair {
	var shortestWait Pair
	minTs := int(^uint(0) >> 1)
	for _, p := range pairs {
		v := Abs(p.Timestamp - ts)
		if v < minTs && p.Timestamp > ts {
			fmt.Printf("p:%#v\n", p)
			minTs = v
			shortestWait = p
		}
	}
	return shortestWait
}

func NextTimestamp(start, step, id, remainder int) (int, int) {
	mult := step * id
	i := start
	for ; i < mult; i += step {
		if (i+remainder)%id == 0 {
			break
		}
	}
	return i, mult
}

func FindCommonTimestamp(buses []*Bus) int {
	sort.Slice(buses, func(i, j int) bool {
		return buses[i].ID < buses[i].ID
	})
	ts := 1
	step := 1
	for _, b := range buses {
		ts, step = NextTimestamp(ts, step, b.ID, b.N)
	}
	return ts
}

func main() {
	filename := os.Args[1]
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	puzzle := LoadPuzzle(f)
	pairs := puzzle.ProcessBuses()
	p := FindBusPair(puzzle.Timestamp, pairs)

	fmt.Printf("Result: %d\n", p.Bus.ID*(p.Timestamp-puzzle.Timestamp))
	fmt.Printf("Result: %d\n", FindCommonTimestamp(puzzle.Buses))
}
