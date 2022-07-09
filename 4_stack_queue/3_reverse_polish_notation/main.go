// Stackを使用して、中置記法 (Infix Notation) の式を逆ポーランド記法 (Reverse Polish Notation)へ変換するプログラムを実装しましょう。
// 【条件】
// 式は、1桁の正の数（1~9）と4つの演算子（+, -, *, /）で構成されます。
// 標準入力から中置記法の式を読み取り、それを逆ポーランド記法に変換し、出力します。
// 変換した式を計算し、計算結果を出力します。
// 【例】
// 入力 : 3+54
// 出力 : 354+ = 23

package main

import (
	"fmt"
	"strconv"
)

// 特に置き換える必要はありませんが、汎用性のため Element に置き換えています
type Element string

type Stack struct {
	data []Element
}

type StackI interface {
	isEmpty() bool
	Push(operator Element)
	Pop() Element
	Top() Element
}

var _ StackI = &Stack{}

func (stack *Stack) isEmpty() bool {
	return len(stack.data) == 0
}

func (stack *Stack) Push(operator Element) {
	stack.data = append(stack.data, operator)
}

func (stack *Stack) Pop() Element {
	top := stack.Top()
	stack.data = stack.data[:len(stack.data)-1]
	return top
}

func (stack *Stack) Top() Element {
	if stack.isEmpty() {
		return ""
	}
	return stack.data[len(stack.data)-1]
}

// 演算子であるかどうかの判定
func isOperator(s string) bool {
	operators := []string{"+", "-", "*", "/"}
	for _, o := range operators {
		if s == o {
			return true
		}
	}
	return false
}

func inStackPriority(s string) int {
	switch s {
	case "+", "-":
		return 1
	case "*", "/":
		return 2
	default:
		return 0
	}
}

func inComingPriority(s string) int {
	switch s {
	case "+", "-":
		return 1
	case "*", "/":
		return 2
	default:
		return 0
	}
}

func convertRPN(input string) string {
	var stack *Stack = &Stack{} // Stack の初期化, 演算子を Push/Pop していく
	var expression string       // 出力する式
	for _, v := range input {
		s := fmt.Sprintf("%c", v)
		if isOperator(s) { // 演算子のとき
			for {
				if inComingPriority(s) > inStackPriority(string(stack.Top())) {
					stack.Push(Element(s))
					break
				} else {
					top := stack.Pop()
					expression += string(top)
				}
			}
		} else { // 正数のとき
			expression += s
		}
	}

	for {
		if stack.isEmpty() {
			break
		} else {
			top := stack.Pop()
			expression += string(top)
		}
	}

	return expression
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// 計算
func operate(s string, a, b int) int {
	switch s {
	case "+":
		return b + a
	case "-":
		return b - a
	case "*":
		return b * a
	case "/":
		return b / a
	default:
		return 0
	}
}

func calculateRPN(expRPN string) int {
	var stack *Stack = &Stack{} // Stack の初期化, 正数を Push/Pop していく
	for _, v := range expRPN {
		s := fmt.Sprintf("%c", v)
		if isOperator(s) { // 演算子のとき
			a, err := strconv.Atoi(string(stack.Pop()))
			check(err)
			b, err := strconv.Atoi(string(stack.Pop()))
			check(err)
			stack.Push(Element(fmt.Sprintf("%d", operate(s, a, b)))) // 計算結果を Stack に Push
		} else { // 正数のとき
			stack.Push(Element(s))
		}
	}
	n, err := strconv.Atoi(string(stack.Pop()))
	check(err)
	return n
}

func main() {
	var input string
	fmt.Println("計算式を入力してください (1~9,+,-,*,/) 例: 3+5*4")
	fmt.Scanf("%s", &input)

	expRPN := convertRPN(input)
	fmt.Printf(expRPN)
	result := calculateRPN(expRPN)
	fmt.Printf(" = %d", result)
}

// ############# result #############

// 計算式を入力してください (1~9,+,-,*,/) 例: 3+5*4
// 7-3*4/2+1
// 734*2/-1+ = 2
