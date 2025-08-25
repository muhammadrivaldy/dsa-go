package main

type binaryNode struct {
	value int
	lNode *binaryNode
	rNode *binaryNode
}

func setBinaryTree(arr []int) *binaryNode {

	head := &binaryNode{value: arr[0]}
	node := head

	for i := 1; i < len(arr); i++ {
		node.addNode(arr[i])
	}

	return head
}

func (b *binaryNode) addNode(val int) {

	node := b

	if node.value > val {
		if node.lNode == nil {
			node.lNode = &binaryNode{value: val}
			return
		}
		node = node.lNode
	} else {
		if node.rNode == nil {
			node.rNode = &binaryNode{value: val}
			return
		}
		node = node.rNode
	}

	node.addNode(val)
}

func binaryTreeInOrder(node *binaryNode) []int {

	output := []int{}

	if node != nil {

		lArr := binaryTreeInOrder(node.lNode)
		root := node.value
		rArr := binaryTreeInOrder(node.rNode)

		temp := []int{}
		temp = append(temp, lArr...)
		temp = append(temp, root)
		temp = append(temp, rArr...)

		output = temp
	}

	return output
}

func binaryTreePreOrder(node *binaryNode) []int {

	output := []int{}

	if node != nil {

		root := node.value
		lArr := binaryTreePreOrder(node.lNode)
		rArr := binaryTreePreOrder(node.rNode)

		temp := []int{}
		temp = append(temp, root)
		temp = append(temp, lArr...)
		temp = append(temp, rArr...)

		output = temp
	}

	return output
}

func binaryTreePostOrder(node *binaryNode) []int {

	output := []int{}

	if node != nil {

		lArr := binaryTreePostOrder(node.lNode)
		rArr := binaryTreePostOrder(node.rNode)
		root := node.value

		temp := []int{}
		temp = append(temp, lArr...)
		temp = append(temp, rArr...)
		temp = append(temp, root)

		output = temp
	}

	return output
}
