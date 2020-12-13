package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Pair struct {
	Bus       int
	Timestamp int
}

type Puzzle struct {
	Timestamp int
	Buses     []int
}

func LoadPuzzle(input io.ReadCloser) *Puzzle {
	puzzle := &Puzzle{
		Buses: []int{},
	}
	defer input.Close()
	data, _ := ioutil.ReadAll(input)
	txt := string(data)
	lines := strings.Split(txt, "\n")
	ts, _ := strconv.Atoi(lines[0])
	puzzle.Timestamp = ts
	for _, v := range strings.Split(lines[1], ",") {
		if v == "x" {
			continue
		}
		bus, _ := strconv.Atoi(v)
		puzzle.Buses = append(puzzle.Buses, bus)
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
		v := float64(p.Timestamp) / float64(b)
		f := int(math.Floor(v)) * b
		c := int(math.Ceil(v)) * b

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

func main() {
	filename := os.Args[1]
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	puzzle := LoadPuzzle(f)
	pairs := puzzle.ProcessBuses()
	p := FindBusPair(puzzle.Timestamp, pairs)

	fmt.Printf("Result: %d\n", p.Bus*(p.Timestamp-puzzle.Timestamp))
}
