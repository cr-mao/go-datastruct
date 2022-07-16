package leetcode

/**
这里我们巧妙运用了一个dummy节点来做我们的操作，这个方法可以很有效地避免一些边界问题，因此希望大家能够养成一个习惯，只要做到链表类的问题，就可以无脑使用一个dummy节点。
*/

func deleteDuplicates(head *ListNode) *ListNode {
	dump := head
	for head != nil {
		for head.Next != nil && head.Next.Val == head.Val {
			head.Next = head.Next.Next
		}
		head = head.Next
	}
	return dump
}

func deleteDuplicates2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	dummy := &ListNode{0, head}

	cur := dummy
	for cur.Next != nil && cur.Next.Next != nil {
		if cur.Next.Val == cur.Next.Next.Val {
			x := cur.Next.Val
			for cur.Next != nil && cur.Next.Val == x {
				cur.Next = cur.Next.Next
			}
		} else {
			cur = cur.Next
		}
	}

	return dummy.Next
}
