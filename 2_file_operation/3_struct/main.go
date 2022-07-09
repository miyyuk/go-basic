// 次は構造体の読み書きをしてみましょう。
// 名前、電話番号、メールアドレスが記載されたアドレス帳を想定します。
// アドレスが記載されたファイルから内容を読み取り、その中身をコピーしたファイルを作成します。
// fread()でファイルから読み取り、fwrite()でファイルに書き込むプログラムを作成しましょう。

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type Address struct {
	name             string
	telephone_number string
	email_address    string
}

func check(e error) {
	if e != nil {
		if e != io.EOF {
			panic(e)
		}
	}
}

func fread(file string) []*Address {
	f, err := os.Open(file)
	check(err)

	defer f.Close()

	r := bufio.NewReader(f)
	var data []*Address
	for {
		n, err := r.ReadString('\n') // nameの読み取り
		check(err)
		t, err := r.ReadString('\n') // telephone_numberの読み取り
		check(err)
		e, err := r.ReadString('\n') // email_addressの読み取り
		check(err)

		address := &Address{name: n, telephone_number: t, email_address: e}
		data = append(data, address)

		_, err = r.ReadString('\n')
		if err == io.EOF {
			break
		}
		check(err)
	}

	return data
}

func fwrite(file string, data []*Address) {
	f, err := os.Create(file)
	check(err)

	defer f.Close()

	for _, v := range data {
		f.WriteString(fmt.Sprintf("%s%s%s\n", v.name, v.telephone_number, v.email_address))
	}
}

func main() {
	data := fread("address.txt")
	fwrite("addressCopy.txt", data)
}
