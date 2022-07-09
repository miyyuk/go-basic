package main

import "fmt"

func main() {
	word := "hello world"
	counter := map[rune]int{}
	for _, c := range word {
		counter[c]++
	}

	seen := map[rune]bool{}
	for _, c := range word {
		if seen[c] {
			continue
		}
		seen[c] = true
		fmt.Printf("letter: %c count: %d\n", c, counter[c])
	}
}
