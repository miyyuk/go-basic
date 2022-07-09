package main

import "fmt"

func NaiveSearch(txt, pat string) {
	if len(txt) <= len(pat) {
		fmt.Println("text is shorter than pattern!")
	}

	patLastIndex := len(pat) - 1 // pattern の最後の index
	var idx []int                // pattern と一致した text 部分の先頭の index を格納する slice
	for i := 0; i < len(txt)-patLastIndex; i++ {
		for j := 0; j < len(pat); j++ {
			if txt[i+j] != pat[j] {
				break // 文字が一致しないとき、 i を一つ進める
			}
			if j == patLastIndex { // pattern がすべて一致したとき
				idx = append(idx, i)
			}
		}
	}

	if len(idx) != 0 {
		fmt.Printf("match index: %v\n", idx)
	} else {
		fmt.Println("no matching index!")
	}
}

func main() {
	NaiveSearch("xyzzxzyxyxyyzxzyyxyxyyz", "xyxyy")
}

// ############# result #############
// match index: [7 17]
