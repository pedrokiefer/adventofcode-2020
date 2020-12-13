package main

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanLoadNavigationInstructions(t *testing.T) {
	input := ioutil.NopCloser(bytes.NewReader([]byte(`F10
N3
F7
R90
F11`)))

	insts := LoadNavigationInstructions(input)

	assert.Equal(t, 5, len(insts))
	assert.Equal(t, &Instruction{Action: Forward, Value: 10}, insts[0])
}

func TestCanExecuteSingleInstructions(t *testing.T) {
	nav := NewNavigation()

	nav.Execute([]*Instruction{{Action: Forward, Value: 10}})
	assert.Equal(t, Point{X: 10, Y: 0}, nav.CurrentPosition)
	assert.Equal(t, EastHeading, nav.Heading)

	nav.Execute([]*Instruction{{Action: North, Value: 3}})
	assert.Equal(t, Point{X: 10, Y: 3}, nav.CurrentPosition)
	assert.Equal(t, EastHeading, nav.Heading)

	nav.Execute([]*Instruction{{Action: Forward, Value: 7}})
	assert.Equal(t, Point{X: 17, Y: 3}, nav.CurrentPosition)
	assert.Equal(t, EastHeading, nav.Heading)

	nav.Execute([]*Instruction{{Action: Right, Value: 90}})
	assert.Equal(t, Point{X: 17, Y: 3}, nav.CurrentPosition)
	assert.Equal(t, SouthHeading, nav.Heading)

	nav.Execute([]*Instruction{{Action: Forward, Value: 11}})
	assert.Equal(t, Point{X: 17, Y: -8}, nav.CurrentPosition)
	assert.Equal(t, SouthHeading, nav.Heading)
}

func TestCanExecuteInstructions(t *testing.T) {
	input := ioutil.NopCloser(bytes.NewReader([]byte(`F10
N3
F7
R90
F11`)))

	insts := LoadNavigationInstructions(input)

	nav := NewNavigation()
	nav.Execute(insts)

	assert.Equal(t, Point{X: 17, Y: -8}, nav.CurrentPosition)
	assert.Equal(t, SouthHeading, nav.Heading)
	assert.Equal(t, 25, nav.ManhattanDistance())
}

func TestCanExecuteSingleInstructionsWayPoint(t *testing.T) {
	nav := NewNavigation()

	nav.ExecuteWaypoint([]*Instruction{{Action: Forward, Value: 10}})
	assert.Equal(t, Point{X: 100, Y: 10}, nav.CurrentPosition)
	assert.Equal(t, Point{X: 10, Y: 1}, nav.WayPoint)

	nav.ExecuteWaypoint([]*Instruction{{Action: North, Value: 3}})
	assert.Equal(t, Point{X: 100, Y: 10}, nav.CurrentPosition)
	assert.Equal(t, Point{X: 10, Y: 4}, nav.WayPoint)

	nav.ExecuteWaypoint([]*Instruction{{Action: Forward, Value: 7}})
	assert.Equal(t, Point{X: 170, Y: 38}, nav.CurrentPosition)
	assert.Equal(t, Point{X: 10, Y: 4}, nav.WayPoint)

	nav.ExecuteWaypoint([]*Instruction{{Action: Right, Value: 90}})
	assert.Equal(t, Point{X: 170, Y: 38}, nav.CurrentPosition)
	assert.Equal(t, Point{X: 4, Y: -10}, nav.WayPoint)

	nav.ExecuteWaypoint([]*Instruction{{Action: Forward, Value: 11}})
	assert.Equal(t, Point{X: 214, Y: -72}, nav.CurrentPosition)
	assert.Equal(t, Point{X: 4, Y: -10}, nav.WayPoint)
}

func TestCanExecuteInstructionsWaypoint(t *testing.T) {
	input := ioutil.NopCloser(bytes.NewReader([]byte(`F10
N3
F7
R90
F11`)))

	insts := LoadNavigationInstructions(input)

	nav := NewNavigation()
	nav.ExecuteWaypoint(insts)

	assert.Equal(t, Point{X: 214, Y: -72}, nav.CurrentPosition)
	assert.Equal(t, Point{X: 4, Y: -10}, nav.WayPoint)
	assert.Equal(t, 286, nav.ManhattanDistance())
}
