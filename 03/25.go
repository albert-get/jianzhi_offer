package main

import (
	"fmt"
	"harry.com/jianzhi_offer/lianbiao"
)

/**
链表中的数字相加
*/
/**
题目：给定两个表示非负整数的单向链表，请问如何实现这两个整数的相加并且把它们的和仍然用单向链表表示？链表中的每个节点表示整数十进制的一位，并且头节点对应整数的最高位数而尾节点对应整数的个位数。
*/

func addReversed(head1 *lianbiao.ListNode, head2 *lianbiao.ListNode) *lianbiao.ListNode {
	dummy := lianbiao.ListNode{}
	sumNode := &dummy
	carry := 0
	for head1 != nil || head2 != nil {
		sum := 0
		if head1 != nil {
			sum += head1.Val
			head1 = head1.Next
		}
		if head2 != nil {
			sum += head2.Val
			head2 = head2.Next
		}
		sum += carry
		if sum >= 10 {
			carry = 1
			sum = sum - 10
		} else {
			carry = 0
		}
		newNode := &lianbiao.ListNode{Val: sum}
		sumNode.Next = newNode
		sumNode = sumNode.Next
	}
	if carry > 0 {
		sumNode.Next = &lianbiao.ListNode{Val: carry}
	}
	return dummy.Next
}
func addTwoNumbers(head1 *lianbiao.ListNode, head2 *lianbiao.ListNode) *lianbiao.ListNode {
	head1 = lianbiao.ReverseList(head1)
	head2 = lianbiao.ReverseList(head2)
	reversedHead := addReversed(head1, head2)
	return lianbiao.ReverseList(reversedHead)
}
func main() {
	head1 := lianbiao.Constructor([]int{9, 8, 4})
	head2 := lianbiao.Constructor([]int{1, 8})
	result := addTwoNumbers(head1, head2)
	fmt.Println(result)
}
