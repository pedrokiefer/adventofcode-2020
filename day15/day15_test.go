package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanPlayElvesGame(t *testing.T) {
	testInput := []int{0, 3, 6}
	number := PlayElvesGame(testInput, 10)

	assert.Equal(t, 0, number)
}

func TestCanPlayElvesGame2020(t *testing.T) {
	testInput := []int{1, 3, 2}
	number := PlayElvesGame(testInput, 2020)

	assert.Equal(t, 1, number)
}

func TestCanPlayElvesGame20202(t *testing.T) {
	testInput := []int{3, 1, 2}
	number := PlayElvesGame(testInput, 2020)

	assert.Equal(t, 1836, number)
}
