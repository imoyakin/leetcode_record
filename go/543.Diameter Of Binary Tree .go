package leetcode

//二叉树的直径
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//思路： dfs
func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ret := 0
	dfs(root, &ret)
	return ret
}

func dfs(root *TreeNode, ret *int) int {
	if root == nil {
		return 0
	}
	l := dfs(root.Left, ret)
	r := dfs(root.Right, ret)
	*ret = max(*ret, l+r)
	return max(l, r) + 1
}

func max(l, r int) int {
	if l > r {
		return l
	}
	return r
}
