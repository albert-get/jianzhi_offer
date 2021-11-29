package main

import "fmt"

/**
最近最少使用缓存
*/
/**
题目：请设计实现一个最近最少使用（Least Recently Used，LRU）缓存，要求如下两个操作的时间复杂度都是O
（1）。
·get（key）：如果缓存中存在键key，则返回它对应的值；否则返回-1。
·put（key，value）：如果缓存中之前包含键key，则它的值设为value；否则添加键key及对应的值value。在添加一个键时，如果缓存容量已经满了，则在添加新键之前删除最近最少使用的键（缓存中最长时间没有被使用过的元素）。
*/

type ListNode struct {
	key   int
	value int
	next  *ListNode
	prev  *ListNode
}

type LRUCache struct {
	head     *ListNode
	tail     *ListNode
	Map      map[int]*ListNode
	capacity int
}

func Constructor(cap int) LRUCache {
	temp := LRUCache{capacity: cap, head: &ListNode{}, tail: &ListNode{}, Map: make(map[int]*ListNode)}
	temp.head.next = temp.tail
	temp.tail.prev = temp.head
	return temp
}
func (this *LRUCache) get(key int) int {
	node, ok := this.Map[key]
	if !ok {
		return -1
	}
	this.moveToTail(node, node.value)
	return node.value
}
func (this *LRUCache) put(key int, value int) {
	node, ok := this.Map[key]
	if ok {
		this.moveToTail(node, value)
	} else {
		if len(this.Map) == this.capacity {
			toBeDeleted := this.head.next
			this.deleteNode(toBeDeleted)
			delete(this.Map, toBeDeleted.key)
		}
		newNode := &ListNode{key: key, value: value}
		this.insertToTail(newNode)
		this.Map[key] = newNode
	}
}
func (this *LRUCache) deleteNode(node *ListNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}
func (this *LRUCache) insertToTail(node *ListNode) {
	this.tail.prev.next = node
	node.prev = this.tail.prev
	node.next = this.tail
	this.tail.prev = node
}
func (this *LRUCache) moveToTail(node *ListNode, newValue int) {
	this.deleteNode(node)
	node.value = newValue
	this.insertToTail(node)
}
func main() {
	mapwithcache := Constructor(4)
	mapwithcache.put(1, 1)
	mapwithcache.put(2, 2)
	mapwithcache.put(3, 3)
	mapwithcache.put(4, 4)

	mapwithcache.get(2)
	mapwithcache.put(1, 8)
	mapwithcache.put(5, 5)
	fmt.Println(mapwithcache)
}
