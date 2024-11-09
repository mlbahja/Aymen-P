package piscine

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	rplc.Parent = node.Parent
	if node.Parent.Left == node {
		node.Parent.Left = rplc
	} else {
		node.Parent.Right = rplc
	}
	node.Parent = nil
	node.Left = nil
	node.Right = nil
	return root
}
