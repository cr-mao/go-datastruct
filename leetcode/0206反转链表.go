package leetcode

/**
单链表反转
链表中环的检测
两个有序的链表合并
删除链表倒数第 n 个结点
求链表的中间结点
*/

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	arr := make([]*ListNode, 0)
	for head != nil {
		arr = append(arr, head)
		head = head.Next
	}
	newList := arr[len(arr)-1]
	arr = arr[:len(arr)-1]
	tmpList := newList
	for len(arr) > 0 {
		tmpList.Next = arr[len(arr)-1]
		arr = arr[:len(arr)-1]
		tmpList = tmpList.Next
	}
	tmpList.Next = nil
	return newList
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }.

2->1->nil
1->2->3->nil
*/
func reverseList1(head *ListNode) *ListNode {
	var pre *ListNode
	var next *ListNode
	cur := head
	for cur != nil {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}
