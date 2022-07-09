// 事前にread.txtを作成し、適当な言葉を記述しておきます。
// 1. write.txtというファイルを作成しましょう。
// 2. read.txtに記述された言葉を一文字ずつ読み出しましょう。
// 3. read.txtから読み出した文字をwrite.txtに書き出しましょう。

package main

import (
	"fmt"
	"io"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	f, err := os.Open("read.txt") // read.txtを開く
	check(err)

	defer f.Close() // read.txtを閉じる（f.Close()で詳説）

	fw, err := os.Create("write.txt") // write.txtという名前のファイルを作成
	check(err)

	defer fw.Close() // write.txtを閉じる

	b := make([]byte, 5) // 5byteの容量を持ったSlice

	// read.txtの中身を5byteずつ読み取り、読み取った分をwrite.txtに書き込む
	// read.txtに読み取るものがなくなるまで、5byteずつの読み書きを続ける
	for {
		n, err := f.Read(b) // read.txtの中身を5byte分読み取る
		if err != nil {
			if err != io.EOF {
				check(err)
			}
			break // read.txtに読み取るものが無くなったらbreakする(io.EOFで詳説)
		}
		fw.Write(b[:n]) // 読み取った分をwrite.txtに書き込む

		// log
		fmt.Println(n, b[:n])
		fmt.Println(" ", fmt.Sprintf("%c", b[:n]))
	}
}

// >log
// 5 [115 111 109 101 10]
//   [s o m e
// ]
// 5 [119 114 105 116 101]
//   [w r i t e]
// 2 [115 10]
//   [s
// ]
