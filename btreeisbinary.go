package piscine

func minBTree(node *TreeNode) string {
	current := node
	min := current.Data
	for current != nil {
		if min > current.Data {
			min = current.Data
		}
		current = current.Left
	}
	return min
}

func BTreeIsBinary(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left != nil {
		min := minBTree(root.Left)
		if min > root.Data {
			return false
		}
	}
	if root.Right != nil {
		max := maxBTree(root.Right)
		if max < root.Data {
			return false
		}
	}
	return BTreeIsBinary(root.Left) && BTreeIsBinary(root.Right)
}

func maxBTree(node *TreeNode) string {
	current := node
	max := current.Data
	for current != nil {
		if max < current.Data {
			max = current.Data
		}
		current = current.Right
	}
	return max
}
