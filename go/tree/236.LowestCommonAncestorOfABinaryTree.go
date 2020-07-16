package tree

//二叉树的最近公共祖先
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	dp := map[int]*TreeNode{}
	visted := map[int]bool{}

	var dfs func(r *TreeNode)
	dfs = func(r *TreeNode) {
		if r == nil {
			return
		}
		if r.Left != nil {
			dp[r.Left.Val] = r
			dfs(r.Left)
		}
		if r.Right != nil {
			dp[r.Right.Val] = r
			dfs(r.Right)
		}
	}
	dfs(root)
	for p != nil {
		visted[p.Val] = true
		p = dp[p.Val]
	}

	for q != nil {
		if visted[q.Val] == true {
			return q
		}
		q = dp[q.Val]
	}

	return nil
}
