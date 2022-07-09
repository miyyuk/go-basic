// 単語と2つの文字を入力させ、単語に含まれる入力した1つ目の文字を、入力した2つ目の文字に変換して出力するプログラムを作成しましょう。

package main

import (
	"fmt"
)

func main() {
	var word string
	var b, a rune
	fmt.Println("Please input word and letters to convert. ex) papa p m\n> ")
	fmt.Scanf("%s %c %c\n", &word, &b, &a)

	for _, w := range word {
		if w == b {
			w = a
		}
		fmt.Printf("%c", w)
	}
}

// ############# result #############
// Please input word and letters to convert. ex) papa p m
// >
// papa p m

// output >
// mama
