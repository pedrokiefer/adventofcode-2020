package main

import (
	"fmt"
)

func InStack(key int, s []int) int {
	first := -1
	second := -1
	for i := range s {
		p := len(s) - 1 - i
		v := s[p]
		if key == v {
			if first != -1 {
				second = p
				break
			} else {
				first = p
			}
		}
	}
	if first != -1 && second != -1 {
		return first - second
	} else if first != -1 && second == -1 {
		return 0
	} else {
		return -1
	}
}

func PlayElvesGame(input []int, maxTurns int) int {
	m := map[int]int{}
	stack := []int{}
	lastNum := -1
	for i := 0; i < maxTurns; i++ {
		if len(stack) > 0 {
			lastNum = stack[i-1]
		}
		diff := 0
		if v, ok := m[lastNum]; ok {
			diff = i - v
		}
		if lastNum != -1 {
			m[lastNum] = i
		}
		if i < len(input) {
			stack = append(stack, input[i])
		} else {
			stack = append(stack, diff)
		}
		// fmt.Printf("LastNum: %d\n", lastNum)
		// fmt.Printf("m: %#v\n", m)
		// fmt.Printf("stack: %#v\n", stack)
	}
	return stack[maxTurns-1]
}

func main() {
	testInput := []int{1, 0, 16, 5, 17, 4}
	number := PlayElvesGame(testInput, 30000000)
	fmt.Printf("Result: %d\n", number)
}
