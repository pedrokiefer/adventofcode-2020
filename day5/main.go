package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"
)

func ToSeatID(input string) int {
	seatID := uint16(0)
	for i, c := range input {
		if i == 7 {
			seatID = seatID << 3
		}
		if c == 'B' {
			seatID = seatID | (1 << (6 - i))
		}
		if c == 'R' {
			seatID = seatID | (1 << (9 - i))
		}

	}
	return int(seatID)
}

func LoadSeatIDs(input io.ReadCloser) []int {
	seats := []int{}
	s := bufio.NewScanner(input)
	defer input.Close()

	for s.Scan() {
		txt := s.Text()
		if txt == "" {
			continue
		}
		txt = strings.TrimSpace(txt)
		seats = append(seats, ToSeatID(txt))
	}
	sort.Ints(seats)
	return seats
}

func FindEmptySeatID(seats []int) int {
	for i := 1; i < len(seats); i++ {
		if seats[i-1] == seats[i]-2 {
			return seats[i] - 1
		}
	}
	return 0
}

func main() {
	filename := os.Args[1]
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	seats := LoadSeatIDs(f)
	fmt.Printf("SeatID: %d\n", seats[len(seats)-1])
	fmt.Printf("EmptySeat: %d\n", FindEmptySeatID(seats))
}
