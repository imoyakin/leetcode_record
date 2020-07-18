package tree

import "math"

func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt32
	var maxNode func(*TreeNode) int
	maxNode = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := max(maxNode(node.Left), 0)
		right := max(maxNode(node.Right), 0)
		pathMax := left + right + node.Val
		maxSum = max(maxSum, pathMax)
		return node.Val + max(left, right)
	}
	maxNode(root)
	return maxSum
}
