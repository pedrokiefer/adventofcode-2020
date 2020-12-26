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

type MatchPuzzle struct {
	Rules    map[int]string
	Rule0    []string
	Messages []string
}

func Join(a, b []string) []string {
	result := []string{}
	for _, va := range a {
		for _, vb := range b {
			result = append(result, va+vb)
		}
	}
	//fmt.Printf("Join(a: %#v b: %#v) = %#v\n", a, b, result)
	return result
}

func (m *MatchPuzzle) ProcessRule(rule string) []string {
	//fmt.Printf("rule: %s\n", rule)
	if strings.Contains(rule, `"`) {
		return []string{strings.Trim(rule, `"`)}
	}

	match := []string{}
	p := strings.Split(rule, "|")
	if len(p) > 1 {
		for _, v := range p {
			match = append(match, m.ProcessRule(strings.TrimSpace(v))...)
		}
	} else {
		numbers := strings.Split(rule, " ")
		subMatch := []string{}
		empty := true
		for _, v := range numbers {
			x, _ := strconv.Atoi(v)
			r := m.ProcessRule(m.Rules[x])
			if empty {
				subMatch = r
				empty = false
			} else {
				subMatch = Join(subMatch, r)
			}
		}
		match = subMatch
	}
	return match
}

func (m *MatchPuzzle) GenerateRule0() {
	m.Rule0 = m.ProcessRule(m.Rules[0])
}

func (m *MatchPuzzle) MatchRule0() int {
	count := 0
	for _, message := range m.Messages {
		for _, r := range m.Rule0 {
			if message == r {
				count++
			}
		}
	}
	return count
}

func LoadPuzzle(input io.ReadCloser) *MatchPuzzle {
	m := &MatchPuzzle{
		Rules:    map[int]string{},
		Rule0:    []string{},
		Messages: []string{},
	}
	s := bufio.NewScanner(input)
	defer input.Close()
	section := 0
	for s.Scan() {
		txt := strings.TrimSpace(s.Text())
		if txt == "" {
			section++
			continue
		}
		if section == 0 {
			// Rules
			p := strings.SplitN(txt, ":", 2)
			ruleID, _ := strconv.Atoi(p[0])
			m.Rules[ruleID] = strings.TrimSpace(p[1])
		} else if section == 1 {
			// Messages
			m.Messages = append(m.Messages, txt)
		}
	}
	return m
}

func main() {
	filename := os.Args[1]
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	m := LoadPuzzle(f)
	m.GenerateRule0()

	fmt.Printf("Result: %v\n", m.MatchRule0())
	// fmt.Printf("Result: %v\n", precedence)
}
