package main

import "fmt"

func KMPSearch(txt, pat string) {
	if len(txt) < len(pat) {
		fmt.Println("text is shorter than pattern!")
	}

	// スライドする文字数を決める ずらし表 を作成(pattern に対して pattern 自身を比較する)
	table := make([]int, len(pat))
	table[0] = 0
	for i, j := 1, 0; i < len(pat); { // j は pattern 自身の比較で直前までに一致していた文字数
		if pat[i] == pat[j] { // 一致している場合
			j++
			table[i] = j
			i++
		} else { // 一致しない場合
			if j != 0 {
				j = table[j-1]
			} else {
				table[i] = 0
				i++
			}
		}
	}
	fmt.Printf("table: %v\n", table)

	var idx []int // pattern と一致した text 部分の先頭の index を格納する slice
	for i, j := 0, 0; i < len(txt); {
		if txt[i] == pat[j] { // 文字が一致している場合
			j++
			i++
			if j == len(pat) { // pattern がすべて一致したとき
				idx = append(idx, i-j)
				j = table[j-1]
			}
		} else { // 文字が一致しない場合
			if j != 0 {
				j = table[j-1]
			} else {
				i++
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
	KMPSearch("xyzzxzyxyxyyzxzyyxyxyyz", "xyxyy")
}

// ############# result #############
// table: [0 0 1 2 0]
// match index: [7 17]
