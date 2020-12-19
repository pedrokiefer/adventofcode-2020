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

type Op func(a, b int) int

func Sum(a, b int) int {
	return a + b
}

func Multiply(a, b int) int {
	return a * b
}

func ProcessLineRecursive(tokens string) (int, int) {
	left := -1
	var currentOp Op
	for i := 0; i < len(tokens); i++ {
		t := tokens[i]
		if t == '+' {
			currentOp = Sum
		} else if t == '*' {
			currentOp = Multiply
		} else if t == '(' {
			v, advance := ProcessLineRecursive(string(tokens[i+1:]))
			if left == -1 {
				left = v
				i += advance + 1
				continue
			}
			left = currentOp(left, v)
			i += advance + 1
		} else if t == ')' {
			return left, i
		} else if t == ' ' {
			continue
		} else {
			v, _ := strconv.Atoi(string(t))
			if left == -1 {
				left = v
				continue
			}
			left = currentOp(left, v)
		}
	}
	return left, 0
}

func ProcessLineRecursiveWithPrecedence(tokens string) (int, int) {
	left := -1
	multi := []int{}
	var currentOp Op
	for i := 0; i < len(tokens); i++ {
		t := tokens[i]
		if t == '+' {
			currentOp = Sum
		} else if t == '*' {
			multi = append(multi, left)
			left = -1
		} else if t == '(' {
			v, advance := ProcessLineRecursiveWithPrecedence(string(tokens[i+1:]))
			if left == -1 {
				left = v
				i += advance + 1
				continue
			}
			left = currentOp(left, v)
			i += advance + 1
		} else if t == ')' {
			if len(multi) > 0 {
				for _, v := range multi {
					left *= v
				}
			}
			return left, i
		} else if t == ' ' {
			continue
		} else {
			v, _ := strconv.Atoi(string(t))
			if left == -1 {
				left = v
				continue
			}
			left = currentOp(left, v)
		}
	}

	if len(multi) > 0 {
		for _, v := range multi {
			left *= v
		}
	}
	return left, 0
}

func ProcessLine(input string) int {
	r, _ := ProcessLineRecursive(input)
	return r
}

func ProcessLinePrecedence(input string) int {
	r, _ := ProcessLineRecursiveWithPrecedence(input)
	return r
}

func ProcessInput(input io.ReadCloser) (int, int) {
	result := 0
	precedence := 0
	s := bufio.NewScanner(input)
	defer input.Close()
	for s.Scan() {
		txt := strings.TrimSpace(s.Text())
		result += ProcessLine(txt)
		precedence += ProcessLinePrecedence(txt)
	}
	return result, precedence
}

func main() {
	filename := os.Args[1]
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	result, precedence := ProcessInput(f)
	fmt.Printf("Result: %v\n", result)
	fmt.Printf("Result: %v\n", precedence)
}
