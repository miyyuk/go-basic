package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type TreeI interface {
	WordCount(word string)
	PrintTree()
}

type Frequencies struct {
	word  string
	count int
}

type Node struct {
	data  *Frequencies
	left  *Node
	right *Node
}

type Tree struct {
	root *Node
}

var _ TreeI = &Tree{}

func (n *Node) ChildTree() *Tree {
	if n == nil {
		return nil
	}
	return &Tree{root: n}
}

// Tree を形成
func (t *Tree) WordCount(word string) {
	if t.root == nil {
		t.root = &Node{data: &Frequencies{word: word, count: 1}}
	} else {
		if t.root.data.word == word {
			t.root.data.count++
		} else if t.root.data.word < word {
			if rightChild := t.root.right.ChildTree(); rightChild != nil {
				rightChild.WordCount(word)
			} else {
				t.root.right = &Node{data: &Frequencies{word: word, count: 1}}
			}
		} else {
			if leftChild := t.root.left.ChildTree(); leftChild != nil {
				leftChild.WordCount(word)
			} else {
				t.root.left = &Node{data: &Frequencies{word: word, count: 1}}
			}
		}
	}
}

// In-order Traversal
// ソートされた順に単語と出現回数を出力する
func (t *Tree) PrintTree() {
	if t == nil || t.root == nil {
		fmt.Printf("nothing!")
		return
	}
	if t.root.left != nil {
		t.root.left.ChildTree().PrintTree()
	}
	fmt.Printf("%s %d  ", t.root.data.word, t.root.data.count)
	if t.root.right != nil {
		t.root.right.ChildTree().PrintTree()
	}
}

func check(e error) {
	if e != nil {
		if e != io.EOF {
			panic(e)
		}
	}
}

func main() {
	// ファイルから文を読み取って単語を抽出する
	f, err := os.Open("sentence.txt")
	check(err)

	defer f.Close()

	r := bufio.NewReader(f)
	sentence, err := r.ReadString('\n')
	check(err)

	sentence = strings.ToUpper(sentence)  // 読み取った文に含まれれるアルファベットを全て大文字にする
	words := strings.Split(sentence, " ") // 単語ごとに区切って Slice に格納する

	// 単語で Tree を形成する
	tree := &Tree{}
	for _, w := range words {
		tree.WordCount(w)
	}

	// Tree を出力
	tree.PrintTree()
}

// ############# result #############
// A 3  AND 1  BLACK 2  CAT 1  MOUSE 2  SAW 1  SCARED 1  SMALL 1  VERY 2
