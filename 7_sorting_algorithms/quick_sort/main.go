package main

import "fmt"

func QuickSort(numbers []int) []int {
	if len(numbers) < 2 {
		return numbers
	}

	fmt.Printf("\nsorting... %v\n", numbers)

	left, right := 0, len(numbers)-1
	pivot := numbers[right]

	for i := 0; i < right; i++ {
		if numbers[i] < pivot {
			numbers[i], numbers[left] = numbers[left], numbers[i]
			left++

			fmt.Printf("p = %d, left = %d, %v\n", pivot, numbers[left], numbers)
		}
	}
	numbers[left], numbers[right] = numbers[right], numbers[left]

	fmt.Printf("p = %d, left = %d, %v\n", pivot, numbers[left], numbers)

	QuickSort(numbers[:left])
	QuickSort(numbers[left+1:])
	return numbers
}

func main() {
	numbers := []int{7, 65, 8, 32, 4, 21, 10}
	QuickSort(numbers)
}

// ############# result #############

// sorting... [7 65 8 32 4 21 10]
// p = 10, left = 65, [7 65 8 32 4 21 10]
// p = 10, left = 65, [7 8 65 32 4 21 10]
// p = 10, left = 32, [7 8 4 32 65 21 10]
// p = 10, left = 10, [7 8 4 10 65 21 32]

// sorting... [7 8 4]
// p = 4, left = 4, [4 8 7]

// sorting... [8 7]
// p = 7, left = 7, [7 8]

// sorting... [65 21 32]
// p = 32, left = 65, [21 65 32]
// p = 32, left = 32, [21 32 65]
