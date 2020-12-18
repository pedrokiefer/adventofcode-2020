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
	X, Y, Z, W int
}

type Bound struct {
	Lower, Upper int
}

type CubePuzzle struct {
	Positions map[Position]bool
	X         Bound
	Y         Bound
	Z         Bound
	W         Bound
}

func NewCube() CubePuzzle {
	return CubePuzzle{
		Positions: map[Position]bool{},
	}
}

func (c *CubePuzzle) LoadInitialMap(input io.ReadCloser) {
	s := bufio.NewScanner(input)
	defer input.Close()
	x := 0
	c.X.Lower = 0
	c.Z.Lower = 0
	c.Z.Upper = 0
	c.W.Lower = 0
	c.W.Upper = 0
	for s.Scan() {
		txt := s.Text()
		if txt == "" {
			continue
		}
		txt = strings.TrimSpace(txt)
		c.Y.Lower = 0
		for y, v := range txt {
			if v == '#' {
				c.Positions[Position{X: x, Y: y, Z: 0}] = true
			}
			c.Y.Upper = y
		}
		c.X.Upper = x
		x++
	}

}

func (c *CubePuzzle) CountNeighbors(p Position) int {
	count := 0
	for i := (p.X - 1); i <= p.X+1; i++ {
		for j := (p.Y - 1); j <= p.Y+1; j++ {
			for k := (p.Z - 1); k <= p.Z+1; k++ {
				for l := (p.W - 1); l <= p.W+1; l++ {
					n := Position{X: i, Y: j, Z: k, W: l}
					if n == p {
						continue
					}
					//fmt.Printf("n = %#v\n", n)
					if _, ok := c.Positions[n]; ok {
						count++
					}
				}
			}
		}
	}
	//fmt.Printf("p = %#v => %d\n", p, count)
	return count
}

func (c *CubePuzzle) RunCycle() {
	newPositions := map[Position]bool{}

	for i := (c.X.Lower - 1); i <= c.X.Upper+1; i++ {
		for j := (c.Y.Lower - 1); j <= c.Y.Upper+1; j++ {
			for k := (c.Z.Lower - 1); k <= c.Z.Upper+1; k++ {
				for l := (c.W.Lower - 1); l <= c.W.Upper+1; l++ {
					p := Position{X: i, Y: j, Z: k, W: l}
					neighbors := c.CountNeighbors(p)
					//fmt.Printf("n = %d\n", neighbors)
					if _, ok := c.Positions[p]; ok {
						// active
						if neighbors == 2 || neighbors == 3 {
							newPositions[p] = true
						}
					} else {
						// inactive
						if neighbors == 3 {
							newPositions[p] = true
						}
					}
				}
			}
		}
	}
	c.X.Lower--
	c.Y.Lower--
	c.Z.Lower--
	c.W.Lower--
	c.X.Upper++
	c.Y.Upper++
	c.Z.Upper++
	c.W.Upper++
	c.Positions = newPositions
}

func (c *CubePuzzle) Print() {
	for l := c.W.Lower; l < c.W.Upper+1; l++ {
		for k := c.Z.Lower; k < c.Z.Upper+1; k++ {
			fmt.Printf("z=%d, w=%d\n", k, l)
			for i := c.X.Lower; i < c.X.Upper; i++ {
				for j := c.Y.Lower; j < c.Y.Upper; j++ {
					p := Position{X: i, Y: j, Z: k, W: l}
					if _, ok := c.Positions[p]; ok {
						fmt.Printf("#")
					} else {
						// inactive
						fmt.Printf(".")
					}
				}
				fmt.Printf("\n")
			}
			fmt.Printf("\n")
		}
	}
}

func main() {
	filename := os.Args[1]
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	c := NewCube()
	c.LoadInitialMap(f)

	for i := 0; i < 6; i++ {
		c.RunCycle()
		c.Print()
	}

	fmt.Printf("Result: %d\n", len(c.Positions))
}
