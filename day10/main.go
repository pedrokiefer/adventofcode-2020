package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func InputToIntList(input io.ReadCloser) []int {
	results := []int{}
	s := bufio.NewScanner(input)
	defer input.Close()
	for s.Scan() {
		i, err := strconv.Atoi(s.Text())
		if err != nil {
			continue
		}
		results = append(results, i)
	}
	return results
}

func CountDifferences(list []int) []int {
	diffs := make([]int, 3)
	list = SortJolts(list)
	for i := 1; i < len(list); i++ {
		v := list[i] - list[i-1]
		switch v {
		case 1:
			diffs[0]++
		case 2:
			diffs[1]++
		case 3:
			diffs[2]++
		}
	}
	return diffs
}

func SortJolts(input []int) (output []int) {
	sort.Ints(input)
	output = append([]int{0}, input...)
	output = append(output, output[len(output)-1]+3)
	return
}

func CountBranchsRecursive(list []int, level int) int {
	spaces := strings.Repeat("\t", level)
	fmt.Printf("%slist: %#v\n", spaces, list)
	result := 0
	for i := 0; i < len(list); i++ {
		if i+1 >= len(list) {
			result = 1
			break
		}

		first := list[i+1] - list[i]
		if first == 3 {
			continue
		}

		if i+2 > len(list) {
			continue
		}
		second := list[i+2] - list[i]
		if second > 3 {
			continue
		}

		if i+3 > len(list) {
			if second == 2 {
				//fmt.Printf("%ssecond is 2 and no more data: %d\n", spaces, list[i])
				for j := 0; j < 2; j++ {
					result += CountBranchsRecursive(list[i+j+1:], level+1)
				}
				return result
			}
			continue
		}
		third := list[i+3] - list[i]

		if third == 3 {
			fmt.Printf("%sthree branchs * %d == 3\n", spaces, list[i])
			for j := 0; j < 3; j++ {
				result += CountBranchsRecursive(list[i+j+1:], level+1)
			}
			return result
		} else if second == 2 {
			fmt.Printf("%stwo branchs * %d == 2\n", spaces, list[i])
			for j := 0; j < 2; j++ {
				result += CountBranchsRecursive(list[i+j+1:], level+1)
			}
			return result
		}
	}
	return result
}

func CountBranchs(list []int) int {
	branchs := make([]int, len(list))
	for i := len(list) - 1; i >= 0; i-- {
		if i+1 >= len(list) {
			branchs[i] = 1
			continue
		}

		first := list[i+1] - list[i]
		if first == 3 {
			branchs[i] = branchs[i+1]
			continue
		}

		if i+2 > len(list) {
			branchs[i] = branchs[i+1]
			continue
		}
		second := list[i+2] - list[i]
		if second > 3 {
			branchs[i] = branchs[i+1]
			continue
		}

		if i+3 > len(list) {
			if second == 2 {
				branchs[i] = branchs[i+1] + branchs[i+2]
			} else {
				branchs[i] = branchs[i+1]
			}
			continue
		}
		third := list[i+3] - list[i]

		if third == 3 {
			branchs[i] = branchs[i+1] + branchs[i+2] + branchs[i+3]
			continue
		} else if second == 2 {
			branchs[i] = branchs[i+1] + branchs[i+2]
			continue
		}

		branchs[i] = branchs[i+1]
	}
	return branchs[0]
}

func main() {
	filename := os.Args[1]
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	list := InputToIntList(f)
	diffs := CountDifferences(list)

	fmt.Printf("Diffs: %#v\n", diffs)
	fmt.Printf("Product: %d\n", diffs[0]*diffs[2])

	sorted := SortJolts(list)
	result := CountBranchs(sorted)
	fmt.Printf("Combinations: %d\n", result)

}
