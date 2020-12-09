package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func LoadMessage(input io.ReadCloser) []int {
	message := []int{}
	s := bufio.NewScanner(input)
	defer input.Close()

	for s.Scan() {
		n, err := strconv.Atoi(s.Text())
		if err != nil {
			continue
		}
		message = append(message, n)
	}

	return message
}

func ValueInArray(v int, list []int) bool {
	for _, vv := range list {
		if v == vv {
			return true
		}
	}
	return false
}

func comb(slice []int, k int, row []int, result *[][]int) {
	if k == 0 {
		r := make([]int, len(row))
		copy(r, row)
		*result = append(*result, r)
		return
	}
	for i := 0; i < len(slice)+1-k; i++ {
		row[len(row)-k] = slice[i]
		comb(slice[i+1:], k-1, row, result)
	}
	row = nil
}

func Combine(slice []int, k int) [][]int {
	result := [][]int{}
	comb(slice, k, make([]int, k), &result)
	return result
}

func SumCombination(slice []int, k int) []int {
	combination := Combine(slice, k)
	sums := []int{}
	for _, v := range combination {
		sum := v[0] + v[1]
		if !ValueInArray(sum, sums) {
			sums = append(sums, sum)
		}
	}
	return sums
}

func Sum(slice []int) int {
	sum := 0
	for _, v := range slice {
		sum += v
	}
	return sum
}

func SumMinMax(slice []int) int {
	min := int(^uint(0) >> 1)
	max := 0
	for _, v := range slice {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return min + max
}

func FindWrongNumber(msg []int, window int) int {
	for i := 0; i < len(msg); i++ {
		validNums := SumCombination(msg[i:window+i], 2)
		if !ValueInArray(msg[window+i], validNums) {
			return msg[window+i]
		}
	}
	return -1
}

func FindContiguousSum(msg []int, value int) []int {
	for i := 0; i < len(msg); i++ {
		for j := i + 1; j < len(msg); j++ {
			if Sum(msg[i:j]) == value {
				return msg[i:j]
			}
		}
	}
	return msg
}

func main() {
	filename := os.Args[1]
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	msg := LoadMessage(f)
	wrong := FindWrongNumber(msg, 25)

	fmt.Printf("Wrong Number: %d\n", wrong)
	r := FindContiguousSum(msg, wrong)
	s := SumMinMax(r)
	fmt.Printf("Sum MinMax: %d\n", s)
}
