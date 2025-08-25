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
		node = head
	}

	return head
}

func (b *binaryNode) addNode(value int) {

	node := b

	for node != nil {

		if value < node.value {

			if node.lNode == nil {
				node.lNode = &binaryNode{value: value}
				return
			} else {
				node = node.lNode
			}

		} else {

			if node.rNode == nil {
				node.rNode = &binaryNode{value: value}
				return
			} else {
				node = node.rNode
			}
		}
	}
}

func binaryNodeInOrder(node *binaryNode) []int {

	output := []int{}

	if node != nil {

		lArr := binaryNodeInOrder(node.lNode)
		root := node.value
		rArr := binaryNodeInOrder(node.rNode)

		temp := []int{}
		temp = append(temp, lArr...)
		temp = append(temp, root)
		temp = append(temp, rArr...)

		output = append(temp, output...)
	}

	return output
}

func binaryNodePreOrder(node *binaryNode) []int {

	output := []int{}

	if node != nil {

		root := node.value
		lArr := binaryNodePreOrder(node.lNode)
		rArr := binaryNodePreOrder(node.rNode)

		temp := []int{}
		temp = append(temp, root)
		temp = append(temp, lArr...)
		temp = append(temp, rArr...)

		output = append(temp, output...)
	}

	return output
}

func binaryNodePostOrder(node *binaryNode) []int {

	output := []int{}

	if node != nil {

		lArr := binaryNodePostOrder(node.lNode)
		rArr := binaryNodePostOrder(node.rNode)
		root := node.value

		temp := []int{}
		temp = append(temp, lArr...)
		temp = append(temp, rArr...)
		temp = append(temp, root)

		output = append(temp, output...)
	}

	return output
}
