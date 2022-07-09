package main

import "fmt"

func SelectionSort(numbers []int) []int {
	for i := 0; i < len(numbers); i++ {
		min := i
		// ソートされていない部分での最小値の index を探す
		for j := i + 1; j < len(numbers); j++ {
			if numbers[min] > numbers[j] {
				min = j
			}
		}
		// ソートされていない部分の先頭と、ソートされていない部分での最小値を入れ替える
		numbers[i], numbers[min] = numbers[min], numbers[i]

		fmt.Printf("i = %d, %v\n", i, numbers)
	}
	return numbers
}

func main() {
	numbers := []int{7, 65, 8, 32, 4, 21, 10}
	SelectionSort(numbers)
}

// ############# result #############
// i = 0, [4 65 8 32 7 21 10]
// i = 1, [4 7 8 32 65 21 10]
// i = 2, [4 7 8 32 65 21 10]
// i = 3, [4 7 8 10 65 21 32]
// i = 4, [4 7 8 10 21 65 32]
// i = 5, [4 7 8 10 21 32 65]
// i = 6, [4 7 8 10 21 32 65]
