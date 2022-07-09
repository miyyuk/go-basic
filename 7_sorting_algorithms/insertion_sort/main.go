package main

import "fmt"

func InsertionSort(numbers []int) []int {
	for i := 0; i < len(numbers); i++ {
		next := numbers[i] // ソートされていない部分の先頭

		fmt.Printf("next: %d\n", next)

		var j int
		// ソートされている部分で、 next が挿入されるところまで要素をずらしていく
		for j = i - 1; j >= 0 && next < numbers[j]; j-- {
			numbers[j+1] = numbers[j]

			fmt.Printf("j = %d, sorting... %v\n", j, numbers)
		}
		numbers[j+1] = next // next を挿入

		fmt.Printf("after insert next: %v\n\n", numbers)
	}
	return numbers
}

func main() {
	numbers := []int{7, 65, 8, 32, 4, 21, 10}
	InsertionSort(numbers)
}

// ############# result #############
// next: 7
// after insert next: [7 65 8 32 4 21 10]

// next: 65
// after insert next: [7 65 8 32 4 21 10]

// next: 8
// j = 1, sorting... [7 65 65 32 4 21 10]
// after insert next: [7 8 65 32 4 21 10]

// next: 32
// j = 2, sorting... [7 8 65 65 4 21 10]
// after insert next: [7 8 32 65 4 21 10]

// next: 4
// j = 3, sorting... [7 8 32 65 65 21 10]
// j = 2, sorting... [7 8 32 32 65 21 10]
// j = 1, sorting... [7 8 8 32 65 21 10]
// j = 0, sorting... [7 7 8 32 65 21 10]
// after insert next: [4 7 8 32 65 21 10]

// next: 21
// j = 4, sorting... [4 7 8 32 65 65 10]
// j = 3, sorting... [4 7 8 32 32 65 10]
// after insert next: [4 7 8 21 32 65 10]

// next: 10
// j = 5, sorting... [4 7 8 21 32 65 65]
// j = 4, sorting... [4 7 8 21 32 32 65]
// j = 3, sorting... [4 7 8 21 21 32 65]
// after insert next: [4 7 8 10 21 32 65]
