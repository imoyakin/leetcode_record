package leetcode

//翻转二叉树
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	invertTree(root.Left)
	invertTree(root.Right)
	if !(root.Left == nil && root.Right == nil) {
		root.Left, root.Right = root.Right, root.Left
	}
	return root
}
