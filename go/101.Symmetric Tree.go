//对称二叉树
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func isSymmetric(root *TreeNode) bool {
    if root == nil {return true}
    return Symmetric(root.Left, root.Right)
}

func Symmetric(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {return true}
	if left != nil && right == nil {return false}
    if left == nil && right != nil {return false}
	return  left.Val == right.Val && 
	Symmetric(left.Left, right.Right) && 
	Symmetric(left.Right, right.Left)
}