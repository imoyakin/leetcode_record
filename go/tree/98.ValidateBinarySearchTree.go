package tree

import "math"

func isValidBST(root *TreeNode) bool {
	return recursive(root, math.MinInt64, math.MaxInt64)
}

func recursive(root *TreeNode, lower, upper int) bool {
	if root == nil {
		return true
	}
	if root.Val <= lower || root.Val >= upper {
		return false
	}
	return recursive(root.Left, lower, root.Val) && recursive(root.Right, root.Val, upper)
}
