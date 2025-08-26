package main

func depthFirstSearch(node *binaryNode, p, q int) *binaryNode {

	if node == nil {
		return node
	} else {
		val := node.value
		if p == val || q == val {
			return node
		}
	}

	lNode := depthFirstSearch(node.lNode, p, q)
	rNode := depthFirstSearch(node.rNode, p, q)

	if lNode != nil && rNode != nil {
		return node
	}

	if lNode != nil {
		return lNode
	}

	return rNode
}
