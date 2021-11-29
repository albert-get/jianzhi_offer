package main

import (
	"fmt"
	"harry.com/jianzhi_offer/lianbiao"
)

/**
排序的循环链表
*/
/**
问题：在一个循环链表中节点的值递增排序，请设计一个算法在该循环链表中插入节点，并保证插入节点之后的循环链表仍然是排序的。
*/

func insertCore(head *lianbiao.ListNode, node *lianbiao.ListNode) {
	cur := head
	next := head.Next
	biggest := head
	for !(cur.Val <= node.Val && next.Val >= node.Val) && next != head {
		cur = next
		next = next.Next
		if cur.Val >= biggest.Val {
			biggest = cur
		}
	}
	if cur.Val <= node.Val && next.Val >= node.Val {
		cur.Next = node
		node.Next = next
	} else {
		node.Next = biggest.Next
		biggest.Next = node
	}
}
func insert(head *lianbiao.ListNode, insertVal int) *lianbiao.ListNode {
	node := &lianbiao.ListNode{Val: insertVal}
	if head == nil {
		head = node
		head.Next = head
	} else if head.Next == head {
		head.Next = node
		node.Next = head
	} else {
		insertCore(head, node)
	}
	return head
}

func main() {
	head := lianbiao.ConstructorHuan([]int{1, 2, 3, 4}, 0)
	insert(head, 9)
	fmt.Println(head)
}
