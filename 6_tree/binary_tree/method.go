package binary_tree

type Element int

type Node struct {
	left  *Node
	right *Node
}

type Tree struct {
	root *Node
}

func (t *Tree) IsEmptyTree() bool {
	return t.root == nil
}

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

func (t *Tree) CountLeafs() int {
	if t == nil {
		return 0
	}
	if IsLeaf(t.root) {
		return 1
	}
	return t.root.left.ChildTree().CountLeafs() + t.root.right.ChildTree().CountLeafs()
}

func (t *Tree) CountNonLeafNodes() int {
	if t == nil {
		return 0
	}
	if IsLeaf(t.root) {
		return 0
	}
	return 1 + t.root.left.ChildTree().CountNonLeafNodes() + t.root.right.ChildTree().CountNonLeafNodes()
}
