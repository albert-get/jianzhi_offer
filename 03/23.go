package main

import (
	"fmt"
	"harry.com/jianzhi_offer/lianbiao"
)

/**
两个链表的第1个重合节点
*/
/**
题目：输入两个单向链表，请问如何找出它们的第1个重合节点。
*/

func getIntersectionNode(headA *lianbiao.ListNode, headB *lianbiao.ListNode) *lianbiao.ListNode {
	count1 := countList(headA)
	count2 := countList(headB)
	var delta int
	var longer *lianbiao.ListNode
	var shorter *lianbiao.ListNode
	if count1 > count2 {
		delta = count1 - count2
		longer = headA
		shorter = headB
	} else {
		delta = count2 - count1
		longer = headB
		shorter = headA
	}
	node1 := longer
	for i := 0; i < delta; i++ {
		node1 = node1.Next
	}
	node2 := shorter
	for node1 != node2 {
		node1 = node1.Next
		node2 = node2.Next
	}
	return node1
}
func countList(head *lianbiao.ListNode) int {
	count := 0
	for head != nil {
		count++
		head = head.Next
	}
	return count
}
func main() {
	h1, h2 := lianbiao.ConstructorJiao([]int{1, 2, 3}, []int{7, 8}, []int{4, 5, 6})
	fmt.Println(h1,h2)
	fmt.Println(getIntersectionNode(h1,h2))
}
