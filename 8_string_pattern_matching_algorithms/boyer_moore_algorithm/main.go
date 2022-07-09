package main

import "fmt"

func BMSearch(txt, pat string) {
	if len(txt) < len(pat) {
		fmt.Println("text is shorter than pattern!")
	}

	patLastIndex := len(pat) - 1 // pattern の最後の index

	// スライドする文字数を決める ずらし表 を作成
	table := make(map[byte]int)
	for i, j := patLastIndex, 0; i >= 0; i, j = i-1, j+1 { // j は pattern の末尾からの文字数
		if _, ok := table[pat[i]]; !ok { // pattern に含まれる文字でまだ table に存在しない場合
			table[pat[i]] = j
		}
	}
	fmt.Printf("table: %v\n", table)
	for key, value := range table {
		fmt.Printf("key: %s value: %v\n", string(key), value)
	}

	var idx []int // pattern と一致した text 部分の先頭の index を格納する slice
	for i := patLastIndex; i < len(txt); {
		// pattern の後方から探索を行う
		for j := 0; j < len(pat); j++ { // j は pattern の末尾からの文字数
			// 文字が一致している場合
			if txt[i-j] == pat[patLastIndex-j] {
				if j == patLastIndex { // pattern がすべて一致したとき
					idx = append(idx, i-j)
				} else {
					continue
				}
			}

			// 文字が一致しない場合 または pattern がすべて一致した場合、比較位置をずらす
			slide, ok := table[txt[i-j]] // スライドする文字数を table から取得
			if !ok {                     // 比較している text の文字が pattern の中にない場合
				slide = len(pat)
			}
			// スライドする文字数を求める
			if slide-j > 0 {
				i += slide - j
			} else { // 現在の位置より前にスライドしてしまう場合
				i++ // 1 だけスライドする
			}
			break // j++ しても文字が一致することはないので break
		}
	}

	if len(idx) != 0 {
		fmt.Printf("match index: %v\n", idx)
	} else {
		fmt.Println("no matching index!")
	}
}

func main() {
	BMSearch("xyzzxzyxyxyyzxzyyxyxyyz", "xyxyy")
}

// ############# result #############
// table: map[120:2 121:0]
// key: y value: 0
// key: x value: 2
// match index: [7 17]
