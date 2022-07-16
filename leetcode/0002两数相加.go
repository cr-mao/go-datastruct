package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var plus int
	var head, cur *ListNode
	for l1 != nil || l2 != nil {
		a, b := 0, 0
		if l1 != nil {
			a = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			b = l2.Val
			l2 = l2.Next
		}
		sum := a + b + plus
		sum, plus = sum%10, sum/10
		if head == nil {
			head = &ListNode{Val: sum}
			cur = head
		} else {
			cur.Next = &ListNode{Val: sum}
			cur = cur.Next
		}
	}
	if plus > 0 {
		cur.Next = &ListNode{Val: plus}
	}
	return head

}
