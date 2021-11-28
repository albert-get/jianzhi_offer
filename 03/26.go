package main

import (
	"fmt"
	"harry.com/jianzhi_offer/lianbiao"
)

/**
重排链表
 */
/**
问题：给定一个链表，链表中节点的顺序是L0→L1→L2→…→Ln-1→Ln，请问如何重排链表使节点的顺序变成L0→Ln→L1→Ln-1→L2→Ln-2→…？
 */

func reorderList(head *lianbiao.ListNode)  {
	dummy:=&lianbiao.ListNode{}
	dummy.Next=head
	fast:=dummy
	slow:=dummy
	for fast!=nil&&fast.Next!=nil{
		slow=slow.Next
		fast=fast.Next
		if fast.Next!=nil{
			fast=fast.Next
		}
	}
	temp:=slow.Next
	slow.Next=nil
	link(head,lianbiao.ReverseList(temp),dummy)
}
func link(node1 *lianbiao.ListNode,node2 *lianbiao.ListNode,head *lianbiao.ListNode)  {
	prev:=head
	for node1!=nil&&node2!=nil{
		temp:=node1.Next
		prev.Next=node1
		node1.Next=node2
		prev=node2
		node1=temp
		node2=node2.Next
	}
	if node1!=nil{
		prev.Next=node1
	}
}
func main() {
	head:=lianbiao.Constructor([]int{1,2,3,4,5,6,7})
	fmt.Println(head)
	reorderList(head)
	fmt.Println(head)

}
