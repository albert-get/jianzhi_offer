package main

import (
	"fmt"
	"harry.com/jianzhi_offer/lianbiao"
)

/**
回文链表
*/
/**
问题：如何判断一个链表是不是回文？要求解法的时间复杂度是O（n），并且不得使用超过O（1）的辅助空间。如果一个链表是回文，那么链表的节点序列从前往后看和从后往前看是相同的。
*/

func isPalindrome3(head *lianbiao.ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	slow := head
	fast := head.Next
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	secondHalf := slow.Next
	if fast.Next != nil {
		secondHalf = slow.Next.Next
	}
	slow.Next = nil
	return equals(secondHalf, lianbiao.ReverseList(head))
}

func equals(head1 *lianbiao.ListNode, head2 *lianbiao.ListNode) bool {
	for head1 != nil && head2 != nil {
		if head1.Val != head2.Val {
			return false
		}
		head1 = head1.Next
		head2 = head2.Next
	}
	return head1 == nil && head2 == nil
}
func main() {
	head := lianbiao.Constructor([]int{1, 3, 3, 2, 1})
	fmt.Println(isPalindrome3(head))
}
