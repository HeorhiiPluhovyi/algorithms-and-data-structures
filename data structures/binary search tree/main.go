package main

import "fmt"

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

func (n *Node) Insert(k int) {
	if n.Key < k {
		// move right
		if n.Right == nil {
			n.Right = &Node{Key: k}
		} else {
			n.Right.Insert(k)
		}
	} else {
		// move left
		if n.Left == nil {
			n.Left = &Node{Key: k}
		} else {
			n.Left.Insert(k)
		}
	}
}

func (n *Node) Search(k int) bool {
	if n == nil {
		return false
	}

	if n.Key < k {
		// move right
		return n.Right.Search(k)
	} else if n.Key > k {
		// move left
		return n.Left.Search(k)
	}

	return true
}

func inorderTraversal(root *Node) []int {
	res := []int{}

	tree(root, &res)

	return res
}

func tree(root *Node, s *[]int) {
	*s = append(*s, root.Key)
	if root.Left != nil {
		tree(root.Left, s)
	}

	if root.Right != nil {
		tree(root.Right, s)
	}
}

func main() {
	tr := &Node{Key: 100}
	tr.Insert(103)
	tr.Insert(102)
	tr.Insert(104)
	fmt.Println(tr.Right.Left)
	fmt.Println(tr.Search(102))

	fmt.Println(inorderTraversal(tr))
}
