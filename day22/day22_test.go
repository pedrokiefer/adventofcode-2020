package main

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanLoadPlayersDecks(t *testing.T) {
	input := ioutil.NopCloser(bytes.NewReader([]byte(`Player 1:
9
2
6
3
1

Player 2:
5
8
4
7
10`)))

	c := NewCombat()
	c.LoadPlayersDecks(input)

	assert.Equal(t, 9, c.Player1[0])
	assert.Equal(t, 5, c.Player2[0])

	c.DoRound()
	assert.Equal(t, 2, c.Player1[0])
	assert.Equal(t, 6, len(c.Player1))

	c.DoRound()
	assert.Equal(t, 6, c.Player1[0])
	assert.Equal(t, 5, len(c.Player1))
}

func TestCanPlayGame(t *testing.T) {
	input := ioutil.NopCloser(bytes.NewReader([]byte(`Player 1:
9
2
6
3
1

Player 2:
5
8
4
7
10`)))

	c := NewCombat()
	c.LoadPlayersDecks(input)
	c.Play()

	assert.Equal(t, 306, c.Score())
}

func TestCanPlayRecursiveGame(t *testing.T) {
	input := ioutil.NopCloser(bytes.NewReader([]byte(`Player 1:
9
2
6
3
1

Player 2:
5
8
4
7
10`)))

	c := NewCombat()
	c.LoadPlayersDecks(input)
	c.PlayRecursive()

	assert.Equal(t, 291, c.Score())
}

func TestCanPlayRecursiveGameEndsLoop(t *testing.T) {
	input := ioutil.NopCloser(bytes.NewReader([]byte(`Player 1:
43
19

Player 2:
2
29
14`)))

	c := NewCombat()
	c.LoadPlayersDecks(input)
	v := c.PlayRecursive()

	assert.Equal(t, 1, v)
}
