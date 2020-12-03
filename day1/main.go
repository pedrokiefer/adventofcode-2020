package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func InputToIntList(input io.ReadCloser) []int64 {
	results := []int64{}
	s := bufio.NewScanner(input)
	defer input.Close()
	for s.Scan() {
		i, err := strconv.ParseInt(s.Text(), 10, 64)
		if err != nil {
			continue
		}
		results = append(results, i)
	}
	return results
}

func Sum2To2020(numbers []int64) (int64, int64) {
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers); j++ {
			if i == j {
				continue
			}
			if numbers[i]+numbers[j] == 2020 {
				return numbers[i], numbers[j]
			}
		}
	}
	return 0, 0
}

func Sum3To2020(numbers []int64) (int64, int64, int64) {
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers); j++ {
			for k := 0; k < len(numbers); k++ {
				if i == j || i == k || j == k {
					continue
				}
				if numbers[i]+numbers[j]+numbers[k] == 2020 {
					return numbers[i], numbers[j], numbers[k]
				}
			}
		}
	}
	return 0, 0, 0
}

func main() {
	filename := os.Args[1]
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	input := InputToIntList(f)
	a, b := Sum2To2020(input)
	fmt.Printf("Result: %v\n", a*b)
	a, b, c := Sum3To2020(input)
	fmt.Printf("Result: %v\n", a*b*c)
}
