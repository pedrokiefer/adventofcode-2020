package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type Position struct {
	X, Y int
}

type Slope struct {
	Right int
	Down  int
}

type TreeMap struct {
	Lines       []string
	StartWidth  int
	StartHeight int
	P           Position
}

func (tm *TreeMap) Height() int {
	return tm.StartHeight
}

func (tm *TreeMap) Reset() {
	tm.P.X = 0
	tm.P.Y = 0
}

func (tm *TreeMap) Traverse(s *Slope) (string, error) {
	tm.P.X = (tm.P.X + s.Right) % tm.StartWidth
	tm.P.Y = tm.P.Y + s.Down

	if tm.P.Y >= tm.StartHeight {
		return "", fmt.Errorf("Max Height")
	}

	c := string(tm.Lines[tm.P.Y][tm.P.X])

	if c == "." {
		return "O", nil
	} else if c == "#" {
		return "X", nil
	}

	return "", fmt.Errorf("Not match for char %v", c)
}

func LoadTreeMap(input io.ReadCloser) *TreeMap {
	tm := TreeMap{
		Lines:      []string{},
		StartWidth: 0,
		P:          Position{X: 0, Y: 0},
	}
	s := bufio.NewScanner(input)
	defer input.Close()
	for s.Scan() {
		txt := s.Text()
		if txt == "" {
			continue
		}
		txt = strings.TrimSpace(txt)
		tm.Lines = append(tm.Lines, txt)
	}
	tm.StartWidth = len(tm.Lines[0])
	tm.StartHeight = len(tm.Lines)
	return &tm
}

func CountTrees(tm *TreeMap, s *Slope) int {
	count := 0
	for {
		t, err := tm.Traverse(s)
		if err != nil {
			break
		}
		if t == "X" {
			count++
		}
	}
	return count
}

func CountSlopes(tm *TreeMap, slopes []*Slope) int {
	c := 1
	for _, s := range slopes {
		tm.Reset()
		c = c * CountTrees(tm, s)
	}
	return c
}

func main() {
	filename := os.Args[1]
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	m := LoadTreeMap(f)
	c := CountTrees(m, &Slope{Right: 3, Down: 1})

	fmt.Printf("Result: %d\n", c)

	slopes := []*Slope{
		{Right: 1, Down: 1},
		{Right: 3, Down: 1},
		{Right: 5, Down: 1},
		{Right: 7, Down: 1},
		{Right: 1, Down: 2},
	}

	c = CountSlopes(m, slopes)
	fmt.Printf("Result: %d\n", c)

}
