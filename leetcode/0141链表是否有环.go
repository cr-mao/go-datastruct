package leetcode 

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow := head
	quick := head.Next
	for slow != nil && quick != nil && quick.Next != nil {
		slow = slow.Next
		quick = quick.Next.Next
		if slow == quick {
			return true
		}
	}
	return false
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle1(head *ListNode) bool {
	var a = make(map[*ListNode]struct{})
	for head != nil {
		if _, ok := a[head]; ok {
			return true
		}
		a[head] = struct{}{}
		head = head.Next
	}
	return false
}