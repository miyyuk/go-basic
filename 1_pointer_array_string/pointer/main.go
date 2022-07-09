// 3つの変数a,b,cについて、a=b, b=c, c=aとなるような関数をポインタで実装しましょう。

package main

import "fmt"

func main() {
	a, b, c := 1, 2, 3
	swap1(a, b, c)
	fmt.Println(a, b, c) // 1 2 3

	swap2(&a, &b, &c)
	fmt.Println(a, b, c) // 2 3 1
}

func swap1(a, b, c int) {
	a, b, c = b, c, a
}

func swap2(a, b, c *int) {
	*a, *b, *c = *b, *c, *a
}

// ############# result #############
// 1 2 3
// 2 3 1
