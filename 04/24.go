package main

import (
	"fmt"
	"harry.com/jianzhi_offer/lianbiao"
)

/**
反转链表
*/
/**
题目：定义一个函数，输入一个链表的头节点，反转该链表并输出反转后链表的头节点。
*/

func reverseList(head *lianbiao.ListNode) *lianbiao.ListNode {
	//prev := lianbiao.ListNode{}.Next
	var prev *lianbiao.ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}
func main() {
	head := lianbiao.Constructor([]int{1, 2, 3, 4, 5})
	fan := reverseList(head)
	fmt.Println(fan)
	var prev *lianbiao.ListNode
	fmt.Println(prev)
}
