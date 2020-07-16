package tree

//二叉树展开为链表
func flatten(root *TreeNode) {
	for root != nil {
		if root.Left != nil {
			pre := root.Left
			for pre.Right != nil {
				pre = pre.Right
			}
			pre.Right = root.Right
			root.Right = root.Left
			root.Left = nil
			root = root.Right
		} else {
			root = root.Right
		}
	}
}
