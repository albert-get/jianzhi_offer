package main

import (
	"fmt"
	"harry.com/jianzhi_offer/lianbiao"
)

/**
展平多级双向链表
*/
/**
问题：在一个多级双向链表中，节点除了有两个指针分别指向前后两个节点，还有一个指针指向它的子链表，并且子链表也是一个双向链表，它的节点也有指向子链表的指针。请将这样的多级双向链表展平成普通的双向链表，即所有节点都没有子链表。
*/

func flattenGetTail(head *lianbiao.Node) *lianbiao.Node {
	node:=head
	var tail *lianbiao.Node
	for node!=nil{
		next:=node.Next
		if node.Child!=nil{
			child:=node.Child
			childTail:=flattenGetTail(node.Child)
			node.Child=nil
			node.Next=child
			child.Prev=node
			childTail.Next=next
			if next!=nil{
				next.Prev=childTail
			}
			tail=childTail
		}else {
			tail=node
		}
		node=next
	}
	return tail
}
func flatten(head *lianbiao.Node) *lianbiao.Node {
	flattenGetTail(head)
	return head
}

func main() {
	p1 := lianbiao.ToST{Arr: []int{1, 2, 3, 4}, Num: 1}
	p2 := lianbiao.ToST{Arr: []int{5, 6, 7}, Num: 1}
	p3 := lianbiao.ToST{Arr: []int{8, 9}}
	head:=lianbiao.ConstructorST([]lianbiao.ToST{p1, p2, p3})
	flatten(head)
	fmt.Println(head)
}
