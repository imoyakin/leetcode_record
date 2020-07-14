package leetcode

// 路径总和 III
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, sum int) int {
	storage := make(map[int]int)
	storage[0] = 1
	return backtrack(root, 0, sum, storage)
}

func backtrack(root *TreeNode, cur, sum int, s map[int]int) int {
	if root == nil {
		return 0
	}
	cur += root.Val
	var cnt int
	if v, ok := s[cur-sum]; ok {
		cnt = v
	}
	s[cur]++
	cnt += backtrack(root.Left, cur, sum, s)
	cnt += backtrack(root.Right, cur, sum, s)
	s[cur]--
	return cnt
}
