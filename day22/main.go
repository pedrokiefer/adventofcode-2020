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

type Combat struct {
	Player1 []int
	Player2 []int
}

func NewCombat() Combat {
	return Combat{
		Player1: []int{},
		Player2: []int{},
	}
}

func (c *Combat) LoadPlayersDecks(input io.ReadCloser) {
	s := bufio.NewScanner(input)
	defer input.Close()
	curPlayer := &c.Player1
	for s.Scan() {
		txt := strings.TrimSpace(s.Text())
		if txt == "" {
			curPlayer = &c.Player2
			continue
		}
		if v, err := strconv.Atoi(txt); err != nil {
			continue
		} else {
			*curPlayer = append(*curPlayer, v)
		}
	}
}

func (c *Combat) AddCards(winner, loser *[]int) {
	v1 := (*winner)[0]
	v2 := (*loser)[0]

	*winner = (*winner)[1:]
	*winner = append(*winner, v1, v2)

	*loser = (*loser)[1:]
}

func (c *Combat) DoRound() {
	// fmt.Printf("Player 1's deck: %v\n", c.Player1)
	// fmt.Printf("Player 2's deck: %v\n", c.Player2)
	// fmt.Printf("Player 1's plays: %v\n", c.Player1[0])
	// fmt.Printf("Player 2's plays: %v\n", c.Player2[0])
	if c.Player1[0] > c.Player2[0] {
		// fmt.Printf("Player 1 wins the round!\n")
		c.AddCards(&c.Player1, &c.Player2)
	} else if c.Player2[0] > c.Player1[0] {
		// fmt.Printf("Player 2 wins the round!\n")
		c.AddCards(&c.Player2, &c.Player1)
	}
}

func (c *Combat) Play() {
	for {
		if len(c.Player1) == 0 || len(c.Player2) == 0 {
			break
		}
		c.DoRound()
	}
}

func (c *Combat) Score() int {
	var deck *[]int
	if len(c.Player1) > 0 {
		deck = &c.Player1
	} else if len(c.Player2) > 0 {
		deck = &c.Player2
	}
	result := 0
	for i := len(*deck); i > 0; i-- {
		result += i * (*deck)[(len(*deck)-i)]
	}
	return result
}

func main() {
	filename := os.Args[1]
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	c := NewCombat()
	c.LoadPlayersDecks(f)
	c.Play()

	fmt.Printf("Score: %d\n", c.Score())
}
