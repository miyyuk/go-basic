package main

import "fmt"

func binarySearch(slice []int, low int, high int, target int, count int) (int, int) {
	if len(slice) < 1 || target < slice[0] || slice[high] < target {
		return -1, count
	}

	mid := (low + high) / 2 // 中央の index を定義
	count++                 // 比較回数をプラス

	if slice[mid] == target {
		return mid, count
	}

	if low >= high {
		return -1, count // これ以上探索できるデータがないので return
	}

	if slice[mid] > target {
		return binarySearch(slice, low, mid-1, target, count)
	}

	if slice[mid] < target {
		return binarySearch(slice, mid+1, high, target, count)
	}

	return -1, count
}

func main() {
	// 0~99 の slice を定義する
	slice := []int{}
	for i := 0; i < 100; i++ {
		slice = append(slice, i)
	}

	// 探索する数字を入力
	var target int
	fmt.Printf("input target number ▶︎ ")
	fmt.Scanf("%d", &target)

	count := 0 // 探索する数字が見つかるまでに、何回比較が行われたかをカウントする変数を定義
	index, count := binarySearch(slice, 0, len(slice)-1, target, count)
	if index < 0 {
		fmt.Printf("target not found")
	} else {
		fmt.Printf("index: %d, count: %d", index, count)
	}
}

// ############# result #############
// input target number ▶︎ 19
// index: 19, count: 7

// input target number ▶︎ 100
// target not found
