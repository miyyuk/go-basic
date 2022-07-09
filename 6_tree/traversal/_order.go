package traversal

import "fmt"

func (t *Tree) PreorderTraversal() {
	if t == nil || t.root == nil {
		fmt.Printf("nothing!")
		return
	}
	fmt.Printf("%s ", t.root.data)
	if t.root.left != nil {
		t.root.left.ChildTree().PreorderTraversal()
	}
	if t.root.right != nil {
		t.root.right.ChildTree().PreorderTraversal()
	}
}

func (t *Tree) InorderTraversal() {
	if t == nil || t.root == nil {
		fmt.Printf("nothing!")
		return
	}
	if t.root.left != nil {
		t.root.left.ChildTree().InorderTraversal()
	}
	fmt.Printf("%s ", t.root.data)
	if t.root.right != nil {
		t.root.right.ChildTree().InorderTraversal()
	}
}

func (t *Tree) PostorderTraversal() {
	if t == nil || t.root == nil {
		fmt.Printf("nothing!")
		return
	}
	if t.root.left != nil {
		t.root.left.ChildTree().PostorderTraversal()
	}
	if t.root.right != nil {
		t.root.right.ChildTree().PostorderTraversal()
	}
	fmt.Printf("%s ", t.root.data)
}

func (t *Tree) LevelorderTraversal(queue []*Node) {
	if t == nil || t.root == nil {
		fmt.Printf("nothing!")
		return
	}

	if len(queue) == 0 {
		queue = append(queue, t.root)
		t.root.ChildTree().LevelorderTraversal(queue)
	} else {
		node := queue[0]
		queue = queue[1:]
		fmt.Printf("%s ", node.data)
		if node.left != nil {
			queue = append(queue, node.left)
		}
		if node.right != nil {
			queue = append(queue, node.right)
		}
		if len(queue) == 0 {
			return
		}
		queue[0].ChildTree().LevelorderTraversal(queue)
	}
}
