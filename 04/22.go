package main

import (
	"fmt"
)

/**
链表中环的入口节点
*/
/**
题目：如果一个链表中包含环，那么应该如何找出环的入口节点？从链表的头节点开始顺着next指针方向进入环的第1个节点为环的入口节点。例如，在如图4.3所示的链表中，环的入口节点是节点3。
*/

type ListNode2 struct {
	Val  int
	Next *ListNode2
}

func getNodeInLoop(head *ListNode2) *ListNode2 {
	if head == nil || head.Next == nil {
		return nil
	}
	slow := head.Next
	fast := slow.Next
	for slow != nil && fast != nil {
		if slow == fast {
			return slow
		}
		slow = slow.Next
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		}
	}
	return nil
}
func detectCycle(head *ListNode2) *ListNode2 {
	inLoop := getNodeInLoop(head)
	if inLoop == nil {
		return nil
	}
	loopCount := 1
	for n := inLoop; n.Next != inLoop; n = n.Next {
		loopCount++
	}
	fast := head
	for i := 0; i < loopCount; i++ {
		fast = fast.Next
	}
	slow := head
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}
func Constructor2(arr []int, n int) *ListNode2 {
	head := ListNode2{}
	p := &head
	h := &ListNode2{}
	for i, v := range arr {

		cur := &ListNode2{Val: v}
		if i == n {
			h = cur
		}
		p.Next = cur
		p = p.Next
	}
	p.Next = h

	return head.Next
}
func detectCycle2(head *ListNode2) *ListNode2 {
	inLoop := getNodeInLoop(head)
	if inLoop == nil {
		return nil
	}
	node := head
	for node != inLoop {
		node = node.Next
		inLoop = inLoop.Next
	}
	return node
}
func main() {
	head := Constructor2([]int{1, 2, 3, 4, 5, 6}, 2)
	result := detectCycle(head)
	//result := detectCycle2(head)
	fmt.Println(result, result.Val)
}
