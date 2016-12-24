package main

type binaryTreeNode struct {
	parent *binaryTreeNode
	left   *binaryTreeNode
	right  *binaryTreeNode
	value  int
}

func (node *binaryTreeNode) add(i int) {

	if i >= node.value {
		if node.right != nil {
			node.right.add(i)
		} else {
			node.right = &binaryTreeNode{
				value:  i,
				parent: node,
			}
		}

		return
	}

	if node.left != nil {
		node.left.add(i)
	} else {
		node.left = &binaryTreeNode{
			value:  i,
			parent: node,
		}
	}
}

func (node *binaryTreeNode) walk(result []int, pos *int) {
	if node.left != nil {
		node.left.walk(result, pos)
	}
	result[*pos] = node.value
	*pos++
	if node.right != nil {
		node.right.walk(result, pos)
	}
}

// SortBinaryTree sortes input slice of integers
// using binary tree algorithm
func SortBinaryTree(input []int) {
	root := &binaryTreeNode{value: input[0]}
	for _, i := range input[1:len(input)] {
		root.add(i)
	}
	pos := 0
	root.walk(input, &pos)

}
