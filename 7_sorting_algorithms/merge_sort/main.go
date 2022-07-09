package main

import "fmt"

func MergeSort(numbers []int) []int {
	if len(numbers) < 2 {
		return numbers
	}
	middle := len(numbers) / 2
	return merge(MergeSort(numbers[:middle]), MergeSort(numbers[middle:]))
}

func merge(left, right []int) []int {
	var merged []int
	fmt.Printf("\nmerge %v %v\n", left, right)

	// 左のリストと右のリストの要素を比較し、小さい方を merged に追加する
	for i := 0; len(left) > 0 && len(right) > 0; i++ {
		if left[0] < right[0] {
			merged = append(merged, left[0])
			left = left[1:]
		} else {
			merged = append(merged, right[0])
			right = right[1:]
		}
		fmt.Printf("i = %d, %v\n", i, merged)
	}

	// リストに残っている要素を merged に追加
	for _, l := range left {
		merged = append(merged, l)
	}
	for _, r := range right {
		merged = append(merged, r)
	}

	fmt.Printf("sorted %v\n", merged)
	return merged
}

func main() {
	numbers := []int{7, 65, 8, 32, 4, 21, 10}
	MergeSort(numbers)
}

// ############# result #############

// merge [65] [8]
// i = 0, [8]
// sorted [8 65]

// merge [7] [8 65]
// i = 0, [7]
// sorted [7 8 65]

// merge [32] [4]
// i = 0, [4]
// sorted [4 32]

// merge [21] [10]
// i = 0, [10]
// sorted [10 21]

// merge [4 32] [10 21]
// i = 0, [4]
// i = 1, [4 10]
// i = 2, [4 10 21]
// sorted [4 10 21 32]

// merge [7 8 65] [4 10 21 32]
// i = 0, [4]
// i = 1, [4 7]
// i = 2, [4 7 8]
// i = 3, [4 7 8 10]
// i = 4, [4 7 8 10 21]
// i = 5, [4 7 8 10 21 32]
// sorted [4 7 8 10 21 32 65]
