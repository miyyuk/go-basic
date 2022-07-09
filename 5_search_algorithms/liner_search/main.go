package main

import (
	"fmt"
)

func linerSearch(numbers []int, target int) int {
	for i, n := range numbers {
		if n == target {
			fmt.Printf("%d is found: %d\n", target, i)
			return i
		}
	}
	return -1
}

func main() {
	numbers := []int{6, 86, 24, 56, 1, 33, 68, 9, 77}

	i := linerSearch(numbers, 86)
	if i < 0 {
		fmt.Println("target not found")
	}

	i = linerSearch(numbers, 9)
	if i < 0 {
		fmt.Println("target not found")
	}

	i = linerSearch(numbers, 10)
	if i < 0 {
		fmt.Println("target not found")
	}
}

// ############# result #############
// 86 is found: 1
// 9 is found: 7
// target not found
