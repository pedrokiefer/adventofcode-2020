package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	FieldMatcher = regexp.MustCompile("(?P<name>.*): (?P<interval_1>[0-9-]*) or (?P<interval_2>[0-9-]*)")
)

type Interval struct {
	Low, High int
}

type Ticket []int

type TicketsPuzzle struct {
	Fields        map[string][]Interval
	ColNames      []string
	MyTicket      Ticket
	NearbyTickets []Ticket
	ValidTickets  []Ticket
}

func (tp *TicketsPuzzle) ValidateNearbyTickets() int {
	validate := func(v int) int {
		valid := []bool{}
		for _, intervals := range tp.Fields {
			for _, ii := range intervals {
				if v < ii.Low || v > ii.High {
					valid = append(valid, false)
				} else {
					valid = append(valid, true)
				}
			}
		}
		c := 0
		for _, t := range valid {
			if t == true {
				c++
			}
		}
		if c > 0 {
			return 0
		}
		return v
	}
	result := 0
	for _, tk := range tp.NearbyTickets {
		tkValue := 0
		for _, v := range tk {
			value := validate(v)
			tkValue += value
		}

		if tkValue == 0 {
			tp.ValidTickets = append(tp.ValidTickets, tk)
			continue
		}
		result += tkValue
	}
	return result
}

func ValidateIntervals(v int, intervals []Interval) bool {
	if (v >= intervals[0].Low && v <= intervals[0].High) || (v >= intervals[1].Low && v <= intervals[1].High) {
		return true
	}
	return false
}

func (tp *TicketsPuzzle) FindColumnsName() {
	colMapping := map[string][]int{}
	for name, intervals := range tp.Fields {
		cols := []int{}
		for col := 0; col < len(tp.MyTicket); col++ {
			valid := true
			for _, tk := range tp.ValidTickets {
				v := tk[col]
				vi := ValidateIntervals(v, intervals)
				if !vi {
					valid = valid && false
				}
			}
			if valid {
				cols = append(cols, col)
			}
		}
		colMapping[name] = cols
	}
	needed := 1
	usedPositions := map[int]bool{}
restart:
	for name, positions := range colMapping {
		if len(positions) == needed {
			p := -1
			for _, pp := range positions {
				if _, ok := usedPositions[pp]; ok {
					continue
				} else {
					p = pp
				}
			}
			tp.ColNames[p] = name
			usedPositions[p] = true
			delete(colMapping, name)
			needed++
		}
	}
	if len(colMapping) != 1 {
		goto restart
	}
	for k := range colMapping {
		for i, v := range tp.ColNames {
			if v == "" {
				tp.ColNames[i] = k
			}
		}
	}
}

func (tp *TicketsPuzzle) MultiplyDeparture() int {
	result := 1
	for i, v := range tp.ColNames {
		if strings.HasPrefix(v, "departure") {
			result = result * tp.MyTicket[i]
		}
	}
	return result
}

func (tp *TicketsPuzzle) PrintValidTickets() {
	for _, tk := range tp.ValidTickets {
		s := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(tk)), ", "), "[]")
		fmt.Printf("%s\n", s)
	}
}

func LoadTicketsPuzzle(input io.ReadCloser) TicketsPuzzle {
	tkts := TicketsPuzzle{
		Fields:        map[string][]Interval{},
		MyTicket:      Ticket{},
		NearbyTickets: []Ticket{},
		ValidTickets:  []Ticket{},
	}
	s := bufio.NewScanner(input)
	defer input.Close()
	section := 0
	skipLine := 0
	for s.Scan() {
		txt := strings.TrimSpace(s.Text())
		if txt == "" {
			skipLine = 0
			section++
			continue
		}
		if section == 0 {
			// Fields
			parts := FieldMatcher.FindStringSubmatch(txt)
			key := parts[1]
			istr1 := parts[2]
			istr2 := parts[3]

			ip := strings.Split(istr1, "-")
			i1min, _ := strconv.Atoi(ip[0])
			i1max, _ := strconv.Atoi(ip[1])

			ip = strings.Split(istr2, "-")
			i2min, _ := strconv.Atoi(ip[0])
			i2max, _ := strconv.Atoi(ip[1])

			tkts.Fields[key] = []Interval{{i1min, i1max}, {i2min, i2max}}
		} else if section == 1 {
			if skipLine == 0 {
				skipLine++
				continue
			}
			parts := strings.Split(txt, ",")
			for _, p := range parts {
				v, _ := strconv.Atoi(p)
				tkts.MyTicket = append(tkts.MyTicket, v)
			}
		} else if section == 2 {
			// Nearby Tickets
			if skipLine == 0 {
				skipLine++
				continue
			}
			tk := Ticket{}
			parts := strings.Split(txt, ",")
			for _, p := range parts {
				v, _ := strconv.Atoi(p)
				tk = append(tk, v)
			}
			tkts.NearbyTickets = append(tkts.NearbyTickets, tk)
		}
	}
	tkts.ValidTickets = append(tkts.ValidTickets, tkts.MyTicket)
	tkts.ColNames = make([]string, len(tkts.MyTicket))
	return tkts
}

func main() {
	filename := os.Args[1]
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	tkts := LoadTicketsPuzzle(f)

	fmt.Printf("Result: %d\n", tkts.ValidateNearbyTickets())
	tkts.FindColumnsName()
	fmt.Printf("Departure Value: %d\n", tkts.MultiplyDeparture())
}
