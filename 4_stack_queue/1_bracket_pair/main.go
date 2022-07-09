// (())(())((()())))()))()
// のような括弧列がすべてペアを作ることができるか（整合性が取れているか）確認するプログラムを実装しましょう。

package main

import "fmt"

// 成立したペア用のstruct
type Pair struct {
	begin int
	end   int
}

type Stack struct {
	data []int
}

type StackI interface {
	isEmpty() bool
	push(int)
	pop() int
}

var _ StackI = &Stack{}

func (s *Stack) isEmpty() bool {
	return len(s.data) == 0
}

func (s *Stack) push(i int) {
	s.data = append(s.data, i)
}

func (s *Stack) pop() int {
	if s.isEmpty() {
		fmt.Println("error")
		return 0
	}
	last := len(s.data) - 1
	j := s.data[last]
	s.data = s.data[:last]
	return j
}

func check(brackets string) bool {
	fmt.Printf("\n\ncheck: %s\n", brackets)
	var s *Stack = &Stack{} // Stackを初期化
	var pairs []Pair        // 成立したペアを格納するSlice

	// 括弧列を一括弧ずつ処理する
	for i, b := range brackets {
		if fmt.Sprintf("%c", b) == "(" { // "(" のとき
			s.push(i) // "("の index を Stack に push
		} else { // ")" のとき
			if s.isEmpty() {
				fmt.Printf("error") // ペアになる "(" が Stack にないので error!
				return false
			}
			j := s.pop()                                  // ペアが成立したので、 Stack から pop
			pairs = append(pairs, Pair{begin: j, end: i}) // 成立したペアの index を Sliceに格納
		}
	}

	// 括弧列の処理が終了したとき、 Stack が空でなければ、ペアが成立しなかった index が残っているということ
	if !s.isEmpty() {
		fmt.Printf("too many (")
		return false
	}

	// 成立したペアを出力(おまけ)
	for _, pair := range pairs {
		fmt.Printf("(%d, %d)", pair.begin, pair.end)
	}

	return true
}

func main() {
	check("((()())())") // (2, 3)(4, 5)(1, 6)(7, 8)(0, 9)
	check("())")        // error
	check("(()")        // too many (
}

// ############# result #############

// check: ((()())())
// (2, 3)(4, 5)(1, 6)(7, 8)(0, 9)

// check: ())
// error

// check: (()
// too many (
