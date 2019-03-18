package main

import "fmt"

type LinkList struct {
	val int
	left, right *LinkList
}

func CreateNode(val int) *LinkList{
	return &LinkList{val:val}
}

func (node *LinkList) print() {
	fmt.Println(node.val)
}

func (node *LinkList) SetNode(val int) {
	node.val = val
}

func (node *LinkList) Traversal() {
	if node.left != nil {
		node.left.Traversal()
	}

	if node.right != nil {
	node.right.Traversal()
	}

	fmt.Println(node.val)
}

func main()  {
	root := LinkList{}

	root.left = &LinkList{2, nil, nil}
	root.right = &LinkList{val: 5}
	root.left.right = new(LinkList)
	root.right.left = CreateNode(6)

	root.left.print()

	root.left.right.SetNode(10)
	root.left.right.print()


	root.Traversal()
}