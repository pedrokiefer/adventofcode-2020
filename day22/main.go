package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type Combat struct {
	GameID   int
	Player1  []int
	Player2  []int
	P1Hashes map[string]bool
	P2Hashes map[string]bool
}

func NewCombat() Combat {
	return Combat{
		GameID:   1,
		Player1:  []int{},
		Player2:  []int{},
		P1Hashes: map[string]bool{},
		P2Hashes: map[string]bool{},
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

func (c *Combat) CheckAlreadyPlayed() bool {
	p1h := sha256.New()
	for _, v := range c.Player1 {
		bs := make([]byte, 1)
		bs[0] = byte(v)
		p1h.Write(bs)
	}
	s1 := hex.EncodeToString(p1h.Sum(nil))

	p2h := sha256.New()
	for _, v := range c.Player1 {
		bs := make([]byte, 1)
		bs[0] = byte(v)
		p2h.Write(bs)
	}
	s2 := hex.EncodeToString(p2h.Sum(nil))

	if _, ok := c.P1Hashes[s1]; ok {
		return true
	}
	c.P1Hashes[s1] = true

	if _, ok := c.P2Hashes[s2]; ok {
		return true
	}
	c.P2Hashes[s2] = true

	return false
}

func (c *Combat) DoRecursiveRound() int {
	// fmt.Printf("Player 1's deck: %v\n", c.Player1)
	// fmt.Printf("Player 2's deck: %v\n", c.Player2)
	// fmt.Printf("Player 1's plays: %v\n", c.Player1[0])
	// fmt.Printf("Player 2's plays: %v\n", c.Player2[0])
	if c.CheckAlreadyPlayed() {
		return -1
	}

	if len(c.Player1) > c.Player1[0] && len(c.Player2) > c.Player2[0] {
		c.GameID++

		c2 := NewCombat()
		c2.GameID = c.GameID
		c2.Player1 = append(c2.Player1, c.Player1[1:c.Player1[0]+1]...)
		c2.Player2 = append(c2.Player2, c.Player2[1:c.Player2[0]+1]...)

		return c2.PlayRecursive()
	}

	if c.Player1[0] > c.Player2[0] {
		return 1
	} else if c.Player2[0] > c.Player1[0] {
		return 2
	}

	return -1
}

func (c *Combat) Play() {
	for {
		if len(c.Player1) == 0 || len(c.Player2) == 0 {
			break
		}
		c.DoRound()
	}
}

func (c *Combat) PlayRecursive() int {
	// fmt.Printf("Playing game %d\n", c.GameID)
	for {
		if len(c.Player1) == 0 || len(c.Player2) == 0 {
			break
		}
		winner := c.DoRecursiveRound()
		if winner == -1 {
			return 1
		} else if winner == 1 {
			c.AddCards(&c.Player1, &c.Player2)
		} else if winner == 2 {
			c.AddCards(&c.Player2, &c.Player1)
		}
	}

	if len(c.Player1) > 0 {
		return 1
	} else if len(c.Player2) > 0 {
		return 2
	}
	return -1
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
		// fmt.Printf("%d %d\n", (*deck)[(len(*deck)-i)], i)
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

	f, err = os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	c2 := NewCombat()
	c2.LoadPlayersDecks(f)
	c2.PlayRecursive()

	fmt.Printf("Score: %d\n", c2.Score())
}
