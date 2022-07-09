// 1. 1~9 がランダムに並んだ配列から Binary Search Tree を実装しましょう。
// 2. node の数を出力しましょう。
// 3. データが 2 の node を探索しましょう。
// 4. データが 2 の node を削除しましょう。
// 5. 最後に全ての node を削除しましょう。
// 2~5を実行後に node を出力しておきましょう

package main

import "fmt"

type TreeI interface {
	CountNodes() int
	PrintTree()
	Search(x Element) *Node
	InsertNode(x Element)
	DeleteNode(x Element) *Node
	RightSubTree(x Element) *Tree
	FindMostLeft() *Node
	DeleteAllNodes() *Node
}

type Element int

type Node struct {
	data  Element
	left  *Node
	right *Node
}

type Tree struct {
	root *Node
}

var _ TreeI = &Tree{}

func IsLeaf(node *Node) bool {
	if node == nil {
		return true
	}
	return node.left == nil && node.right == nil
}

func (n *Node) ChildTree() *Tree {
	if n == nil {
		return nil
	}
	return &Tree{root: n}
}

func (t *Tree) CountNodes() int {
	if t == nil || t.root == nil {
		return 0
	}
	if IsLeaf(t.root) {
		return 1
	}
	return 1 + t.root.left.ChildTree().CountNodes() + t.root.right.ChildTree().CountNodes()
}

func (t *Tree) PrintTree() {
	if t == nil || t.root == nil {
		fmt.Printf("nothing!")
		return
	}
	if t.root.left != nil {
		t.root.left.ChildTree().PrintTree()
	}
	fmt.Printf("%d ", t.root.data)
	if t.root.right != nil {
		t.root.right.ChildTree().PrintTree()
	}
}

func (t *Tree) Search(x Element) *Node {
	if t == nil {
		return nil
	}
	if t.root.data == x {
		return t.root
	}
	if t.root.data < x {
		return t.root.right.ChildTree().Search(x)
	} else {
		return t.root.left.ChildTree().Search(x)
	}
}

func (t *Tree) InsertNode(x Element) {
	node := &Node{data: x}
	if t.root == nil {
		t.root = node
	} else {
		if t.root.data < x {
			if rightChild := t.root.right.ChildTree(); rightChild != nil {
				rightChild.InsertNode(x)
			} else {
				t.root.right = node
			}
		} else {
			if leftChild := t.root.left.ChildTree(); leftChild != nil {
				leftChild.InsertNode(x)
			} else {
				t.root.left = node
			}
		}
	}
}

func (t *Tree) DeleteNode(x Element) *Node {
	if t == nil {
		return nil
	}
	if t.root.data < x { // root が削除する対象の node ではない
		t.root.right = t.root.right.ChildTree().DeleteNode(x)
		return t.root
	} else if t.root.data > x { // root が削除する対象の node ではない
		t.root.left = t.root.left.ChildTree().DeleteNode(x)
		return t.root
	} else { // root が削除する対象の node である
		if IsLeaf(t.root) { // ① 削除する node が leaf (child node をもたない)のとき
			t.root = nil // 対象の node を削除するだけでよい
		} else if t.root.left == nil && t.root.right != nil { // ② 削除する node が child node を 1 つもつとき
			t.root = t.root.right // child node を root に移動させる
		} else if t.root.left != nil && t.root.right == nil { // ② 削除する node が child node を 1 つもつとき
			t.root = t.root.left // child node を root に移動させる
		} else { // ③ 削除する node が　child node を 2 つもつとき
			rightMin := t.RightSubTree(x).FindMostLeft().data // 削除する node の Right sub tree の最小値を root に移動させる
			t.root.data = rightMin
			t.root.right = t.root.right.ChildTree().DeleteNode(rightMin)
		}
		return t.root
	}
}

func (t *Tree) RightSubTree(x Element) *Tree {
	if t == nil {
		return nil
	}
	if t.root.data == x {
		return t.root.right.ChildTree()
	} else if t.root.data < x {
		return t.root.right.ChildTree().RightSubTree(x)
	} else {
		return t.root.left.ChildTree().RightSubTree(x)
	}
}

func (t *Tree) FindMostLeft() *Node {
	if t == nil {
		return nil
	}
	if t.root.left == nil {
		return t.root
	}
	return t.root.left.ChildTree().FindMostLeft()
}

func (t *Tree) DeleteAllNodes() *Node {
	if t != nil {
		t.root.left = t.root.left.ChildTree().DeleteAllNodes()
		t.root.right = t.root.right.ChildTree().DeleteAllNodes()
		t.root = nil
		return t.root
	}
	return nil
}

func main() {
	tree := &Tree{}
	integers := []Element{5, 2, 3, 8, 9, 1, 4, 6, 7}

	for _, v := range integers {
		tree.InsertNode(v)
	}

	fmt.Printf("count: %d\n", tree.CountNodes())
	tree.PrintTree()

	tree.DeleteNode(2)
	fmt.Printf("\nafter delete 2: ")
	tree.PrintTree()

	tree.DeleteAllNodes()
	fmt.Printf("\nafter delete all nodes: ")
	tree.PrintTree()
}

// ############# result #############
// count: 9
// 1 2 3 4 5 6 7 8 9
// after delete 2: 1 3 4 5 6 7 8 9
// after delete all nodes: nothing!
