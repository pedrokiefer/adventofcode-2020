package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
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

type SinCosPair struct {
	Sin int
	Cos int
}

var (
	SinCosTable = map[int]SinCosPair{
		-270: {Sin: 1, Cos: 0},
		-180: {Sin: 0, Cos: -1},
		-90:  {Sin: -1, Cos: 0},
		90:   {Sin: 1, Cos: 0},
		180:  {Sin: 0, Cos: -1},
		270:  {Sin: -1, Cos: 0},
	}
)

func (p *Point) Rotate(angle int) {
	sc := SinCosTable[angle]
	x := sc.Cos*p.X - sc.Sin*p.Y
	y := sc.Sin*p.X + sc.Cos*p.Y
	p.X = x
	p.Y = y
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
	WayPoint        *Point
}

func NewNavigation() *Navigation {
	wp := NewPoint(10, 1)
	return &Navigation{
		Heading:         EastHeading,
		CurrentPosition: NewPoint(0, 0),
		WayPoint:        &wp,
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
			n.WayPoint.Rotate(-ins.Value)
		case Left:
			n.WayPoint.Rotate(ins.Value)
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

	nav = NewNavigation()
	nav.ExecuteWaypoint(insts)
	fmt.Printf("Distance: %#v\n", nav.ManhattanDistance())
}
