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

type PasswordPolicy struct {
	Low  int
	High int
	Char string
}

type PasswordInput struct {
	Policy   PasswordPolicy
	Password string
}

func (p *PasswordPolicy) CheckPassword(input string) bool {
	// count := strings.Count(input, p.Char)
	// if p.Low <= count && count <= p.High {
	// 	return true
	// }
	first := string(input[p.Low-1]) == p.Char
	second := string(input[p.High-1]) == p.Char
	if first != second {
		return true
	}
	return false
}

func InputToPasswords(input io.ReadCloser) []PasswordInput {
	results := []PasswordInput{}
	s := bufio.NewScanner(input)
	defer input.Close()
	for s.Scan() {
		txt := s.Text()
		if txt == "" {
			continue
		}
		p1 := strings.SplitN(txt, ":", 2)
		password := strings.TrimSpace(p1[1])

		p2 := strings.SplitN(p1[0], " ", 2)
		char := strings.TrimSpace(p2[1])

		p3 := strings.SplitN(p2[0], "-", 2)
		low, err := strconv.ParseInt(p3[0], 10, 64)
		if err != nil {
			continue
		}
		high, err := strconv.ParseInt(p3[1], 10, 64)
		if err != nil {
			continue
		}

		results = append(results, PasswordInput{
			Policy: PasswordPolicy{
				Low:  int(low),
				High: int(high),
				Char: char,
			},
			Password: password,
		})
	}
	return results
}

func CountValidPasswords(input []PasswordInput) int {
	count := 0
	for _, i := range input {
		if i.Policy.CheckPassword(i.Password) {
			count++
		}
	}

	return count
}

func main() {
	filename := os.Args[1]
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	input := InputToPasswords(f)
	result := CountValidPasswords(input)
	fmt.Printf("Result: %v\n", result)
}
