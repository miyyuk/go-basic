package main

import "fmt"

func HeapSort(numbers []int) []int {
	// heap を形成する
	fmt.Printf("Build Heap\n")
	for i := len(numbers)/2 - 1; i >= 0; i-- {
		heapify(numbers, len(numbers), i)
	}

	// heap から要素を一つずつ取り出す (ソートしていく)
	fmt.Printf("\nHeap Sort")
	for i := len(numbers) - 1; i > 0; i-- {
		numbers[i], numbers[0] = numbers[0], numbers[i]

		fmt.Printf("\ni = %d, sorting... %v\n", i, numbers)

		heapify(numbers, i, 0)
	}
	return numbers
}

// i 番目の要素を root とする Tree をヒープ化する.
// len は heap size, root は numbers の index
func heapify(numbers []int, len, root int) {
	largest := root // 最大の要素の index を i (root) としておく
	left := 2*root + 1
	right := 2*root + 2

	if left < len && numbers[left] > numbers[largest] { // left child が root より大きい場合
		largest = left
	}
	if right < len && numbers[right] > numbers[largest] { // right child が root より大きい場合
		largest = right
	}
	if largest != root { // 最大の要素が root ではないとき
		numbers[root], numbers[largest] = numbers[largest], numbers[root] // root の要素と最大の要素を入れ替える

		fmt.Printf("(%d %d %d) heapify... %v\n", root, left, right, numbers)

		heapify(numbers, len, largest)
	}
}

func main() {
	numbers := []int{7, 65, 8, 32, 4, 21, 10}
	HeapSort(numbers)
}

// ############# result #############
// Build Heap
// (2 5 6) heapify... [7 65 21 32 4 8 10]
// (0 1 2) heapify... [65 7 21 32 4 8 10]
// (1 3 4) heapify... [65 32 21 7 4 8 10]

// Heap Sort
// i = 6, sorting... [10 32 21 7 4 8 65]
// (0 1 2) heapify... [32 10 21 7 4 8 65]

// i = 5, sorting... [8 10 21 7 4 32 65]
// (0 1 2) heapify... [21 10 8 7 4 32 65]

// i = 4, sorting... [4 10 8 7 21 32 65]
// (0 1 2) heapify... [10 4 8 7 21 32 65]
// (1 3 4) heapify... [10 7 8 4 21 32 65]

// i = 3, sorting... [4 7 8 10 21 32 65]
// (0 1 2) heapify... [8 7 4 10 21 32 65]

// i = 2, sorting... [4 7 8 10 21 32 65]
// (0 1 2) heapify... [7 4 8 10 21 32 65]

// i = 1, sorting... [4 7 8 10 21 32 65]
