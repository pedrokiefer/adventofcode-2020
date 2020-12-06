package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func CountUniqueDeclarations(input string, everyone bool) int {
	unique := map[rune]int{}
	more := 0
	for _, r := range input {
		if r == '\n' {
			more++
			continue
		}
		if r == ' ' || r == '\t' {
			continue
		}
		if c, ok := unique[r]; ok {
			unique[r] = c + 1
		} else {
			unique[r] = 1
		}
	}

	if everyone && more != 0 {
		count := 0
		for _, v := range unique {
			if v > more {
				count++
			}
		}
		return count
	}
	return len(unique)
}

func LoadDeclarations(input io.ReadCloser, everyone bool) int {
	total := 0
	s := bufio.NewScanner(input)
	defer input.Close()

	buf := ""
	for s.Scan() {
		txt := s.Text()
		if txt == "" {
			total += CountUniqueDeclarations(buf, everyone)
			buf = ""
			continue
		}
		txt = strings.TrimSpace(txt)
		if buf == "" {
			buf += txt
		} else {
			buf = buf + "\n" + txt
		}
	}

	if buf != "" {
		total += CountUniqueDeclarations(buf, everyone)
	}

	return total
}

func main() {
	filename := os.Args[1]
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	decs := LoadDeclarations(f, false)
	fmt.Printf("Unique Declarations By Group: %d\n", decs)

	f, err = os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	decs = LoadDeclarations(f, true)
	fmt.Printf("Unique Declarations Global: %d\n", decs)
}
