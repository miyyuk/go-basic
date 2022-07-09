// 任意の英語の文章(すべて小文字)で、各文字が出現する回数を出力するプログラムを作成しましょう。

package main

import "fmt"

func main() {
	word := "hello world"

	var elements []string  // 文字を入れておくSlice
	var elementCount []int // 文字が何回出現したかを入れておくSlice

	for _, w := range word {
		letter := fmt.Sprintf("%c", w)  // w は rune型(後述) なので、 string型 に変換
		if !isExist(letter, elements) { // 初めて出現した文字について実行
			elements = append(elements, letter)
			elementCount = append(elementCount, count(letter, word))
		}
	}

	for i := 0; i < len(elements); i++ {
		fmt.Printf("letter: %s count: %d\n", elements[i], elementCount[i])
	}
}

// 文字が word のなかに何回出現するかカウントする
func count(letter string, word string) int {
	count := 0
	for _, w := range word {
		l := fmt.Sprintf("%c", w)
		if l == letter {
			count++
		}
	}
	return count
}

// すでにカウントした文字かどうかの判定
func isExist(letter string, elements []string) bool {
	for _, e := range elements {
		if letter == e {
			return true
		}
	}
	return false
}

// ############# result #############
// letter: h count: 1
// letter: e count: 1
// letter: l count: 3
// letter: o count: 2
// letter:   count: 1
// letter: w count: 1
// letter: r count: 1
// letter: d count: 1
