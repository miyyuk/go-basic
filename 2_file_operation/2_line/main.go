// 一つ前のExerciseでは一文字ずつ読み取りましたが、次は一行ずつ読み取りましょう。
// ついでに、行数も出力しましょう。

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func check(e error) {
	if e != nil {
		if e != io.EOF {
			panic(e)
		}
	}
}

func main() {
	f, err := os.Open("text.txt")
	check(err)

	defer f.Close()

	r := bufio.NewReader(f) // Readerを初期化する

	var lines []string // 読み取ったStringを格納しておくSlice
	for {
		l, err := r.ReadString('\n') // 区切り文字である改行コード('\n')まで読み取る

		lines = append(lines, l) // 読み取ったStringをSliceに追加
		if err == io.EOF {
			break
		}
		check(err)
	}

	for i, l := range lines {
		fmt.Printf("%d %s", i+1, l)
	}
}

// ############# result #############
// 1 This is sample file.
// 2 Have a nice day!
