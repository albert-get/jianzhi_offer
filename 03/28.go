package main

import "harry.com/jianzhi_offer/lianbiao"

/**
展平多级双向链表
 */
/**
问题：在一个多级双向链表中，节点除了有两个指针分别指向前后两个节点，还有一个指针指向它的子链表，并且子链表也是一个双向链表，它的节点也有指向子链表的指针。请将这样的多级双向链表展平成普通的双向链表，即所有节点都没有子链表。
 */

func main() {
	lianbiao.ConstructorST([][]int{{{{1,2,3},2}}})
}

