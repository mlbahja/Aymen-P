package piscine

func enqueue(queue *([](*TreeNode)), element *TreeNode) {
	*queue = append(*queue, element)
}

func dequeue(queue *([](*TreeNode))) *TreeNode {
	current := (*queue)[0]
	*queue = (*queue)[1:]
	return current
}

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	queue := [](*TreeNode){root}
	if root != nil {
		f(root.Data)
	}
	for len(queue) != 0 {
		current := dequeue(&queue)
		if current.Left != nil {
			f(current.Left.Data)
		}
		if current.Right != nil {
			f(current.Right.Data)
		}
		if current.Left != nil {
			enqueue(&queue, current.Left)
		}
		if current.Right != nil {
			enqueue(&queue, current.Right)
		}
	}
}
