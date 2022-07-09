package main

import "fmt"

func BubbleSort(numbers []int) []int {
	for i := 0; i < len(numbers)-1; i++ {
		for j := 0; j < len(numbers)-i-1; j++ {
			if numbers[j] > numbers[j+1] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}
		}
		fmt.Printf("i = %d, %v\n", i, numbers)
	}
	return numbers
}

func main() {
	numbers := []int{7, 65, 8, 32, 4, 21, 10}
	BubbleSort(numbers)
}

// ############# result #############
// i = 0, [7 8 32 4 21 10 65]
// i = 1, [7 8 4 21 10 32 65]
// i = 2, [7 4 8 10 21 32 65]
// i = 3, [4 7 8 10 21 32 65]
// i = 4, [4 7 8 10 21 32 65]
// i = 5, [4 7 8 10 21 32 65]
