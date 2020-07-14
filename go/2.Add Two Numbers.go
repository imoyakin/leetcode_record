package leetcode

//两数相加
/**
 * Definition for singly-linked list.

 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	ret := new(ListNode)
	cur := ret
	var c int
	for l1 != nil || l2 != nil || c > 0 {
		cur.Next = new(ListNode)
		cur = cur.Next
		if l1 != nil {
			c += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			c += l2.Val
			l2 = l2.Next
		}
		cur.Val = c % 10
		c = c / 10
	}
	return ret.Next
}
