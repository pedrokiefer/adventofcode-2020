package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	X, Y int
}

func NewPoint(x, y int) Point {
	return Point{X: x, Y: y}
}

func (p *Point) Add(x, y int) {
	p.X += x
	p.Y += y
}

func (p *Point) Multiply(v int) (x, y int) {
	x = p.X * v
	y = p.Y * v
	return
}

var (
	NorthHeading    = Point{X: 0, Y: 1}
	EastHeading     = Point{X: 1, Y: 0}
	SouthHeading    = Point{X: 0, Y: -1}
	WestHeading     = Point{X: -1, Y: 0}
	HeadingSequence = []Point{NorthHeading, EastHeading, SouthHeading, WestHeading}
)

func IndexOf(p Point, ps []Point) int {
	for k, v := range ps {
		if p == v {
			return k
		}
	}
	return -1
}

type Action string

const (
	North   = Action("N")
	East    = Action("E")
	South   = Action("S")
	West    = Action("W")
	Forward = Action("F")
	Right   = Action("R")
	Left    = Action("L")
)

type Instruction struct {
	Action Action
	Value  int
}

type Navigation struct {
	Heading         Point
	CurrentPosition Point
	WayPoint        Point
}

func NewNavigation() *Navigation {
	return &Navigation{
		Heading:         EastHeading,
		CurrentPosition: NewPoint(0, 0),
		WayPoint:        NewPoint(10, 1),
	}
}

func LoadNavigationInstructions(input io.ReadCloser) []*Instruction {
	ins := []*Instruction{}
	s := bufio.NewScanner(input)
	defer input.Close()
	for s.Scan() {
		txt := strings.TrimSpace(s.Text())
		a := Action(txt[0])
		txt = string(txt[1:])
		v, _ := strconv.Atoi(txt)
		ins = append(ins, &Instruction{Action: a, Value: v})
	}
	return ins
}

func (n *Navigation) Execute(insts []*Instruction) {
	for _, ins := range insts {
		switch ins.Action {
		case North:
			n.CurrentPosition.Add(0, ins.Value)
		case South:
			n.CurrentPosition.Add(0, -ins.Value)
		case East:
			n.CurrentPosition.Add(ins.Value, 0)
		case West:
			n.CurrentPosition.Add(-ins.Value, 0)
		case Forward:
			x, y := n.Heading.Multiply(ins.Value)
			n.CurrentPosition.Add(x, y)
		case Right:
			hi := IndexOf(n.Heading, HeadingSequence)
			v := (ins.Value / 90)
			v = ((hi+v)%4 + 4) % 4
			n.Heading = HeadingSequence[v]
		case Left:
			hi := IndexOf(n.Heading, HeadingSequence)
			v := (ins.Value / 90)
			v = ((hi-v)%4 + 4) % 4
			n.Heading = HeadingSequence[v]
		}
	}
}

func (n *Navigation) ExecuteWaypoint(insts []*Instruction) {
	for _, ins := range insts {
		switch ins.Action {
		case North:
			n.WayPoint.Add(0, ins.Value)
		case South:
			n.WayPoint.Add(0, -ins.Value)
		case East:
			n.WayPoint.Add(ins.Value, 0)
		case West:
			n.WayPoint.Add(-ins.Value, 0)
		case Forward:
			x, y := n.WayPoint.Multiply(ins.Value)
			n.CurrentPosition.Add(x, y)
		case Right:
			radians := float64(ins.Value) * math.Pi / 180
			x := math.Cos(-radians)*float64(n.WayPoint.X) - math.Sin(-radians)*float64(n.WayPoint.Y)
			y := math.Sin(-radians)*float64(n.WayPoint.X) + math.Cos(-radians)*float64(n.WayPoint.Y)
			n.WayPoint.X = int(x)
			n.WayPoint.Y = int(y)
		case Left:
			radians := float64(ins.Value) * math.Pi / 180
			x := math.Cos(radians)*float64(n.WayPoint.X) - math.Sin(radians)*float64(n.WayPoint.Y)
			y := math.Sin(radians)*float64(n.WayPoint.X) + math.Cos(radians)*float64(n.WayPoint.Y)
			n.WayPoint.X = int(x)
			n.WayPoint.Y = int(y)
		}
	}
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (n *Navigation) ManhattanDistance() int {
	return Abs(n.CurrentPosition.X) + Abs(n.CurrentPosition.Y)
}

func main() {
	filename := os.Args[1]
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	insts := LoadNavigationInstructions(f)
	nav := NewNavigation()
	nav.Execute(insts)

	fmt.Printf("Distance: %#v\n", nav.ManhattanDistance())

	f, err = os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	insts = LoadNavigationInstructions(f)
	nav = NewNavigation()
	nav.ExecuteWaypoint(insts)

	fmt.Printf("Distance: %#v\n", nav.ManhattanDistance())
}
