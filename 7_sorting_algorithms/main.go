// 1~500 の数字で構成される、要素数 500 のランダムな配列を Insertion sort と　Quick sort でソートしましょう。
// その実行速度を比較してみましょう。

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func initList(n int) []int {
	var numbers []int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < n; i++ {
		numbers = append(numbers, rand.Intn(500))
	}
	return numbers
}

func InsertionSort(numbers []int) []int {
	for i := 0; i < len(numbers); i++ {
		next := numbers[i]
		var j int
		for j = i - 1; j >= 0 && next < numbers[j]; j-- {
			numbers[j+1] = numbers[j]
		}
		numbers[j+1] = next
	}
	return numbers
}

func QuickSort(numbers []int) []int {
	if len(numbers) < 2 {
		return numbers
	}
	left, right := 0, len(numbers)-1
	pivot := numbers[right]

	for i := 0; i < right; i++ {
		if numbers[i] < pivot {
			numbers[i], numbers[left] = numbers[left], numbers[i]
			left++
		}
	}
	numbers[left], numbers[right] = numbers[right], numbers[left]

	QuickSort(numbers[:left])
	QuickSort(numbers[left+1:])
	return numbers
}

func main() {
	numbers := initList(500)
	start := time.Now()
	InsertionSort(numbers)
	runtime := time.Since(start)
	fmt.Printf("Insertion sort ▶︎ runtime: %v\n", runtime)

	numbers = initList(500)
	start = time.Now()
	QuickSort(numbers)
	runtime = time.Since(start)
	fmt.Printf("Quick sort ▶︎ runtime: %v\n", runtime)
}

// ############# result #############
// Insertion sort ▶︎ runtime: 44.611µs
// Quick sort ▶︎ runtime: 23.602µs
