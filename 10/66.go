package main

import "fmt"

/**
单词之和
*/
/**
题目：请设计实现一个类型MapSum，它有如下两个操作。
·函数insert，输入一个字符串和一个整数，在数据集合中添加一个字符串及其对应的值。如果数据集合中已经包含该字符串，则将该字符串对应的值替换成新值。
·函数sum，输入一个字符串，返回数据集合中所有以该字符串为前缀的字符串对应的值之和。
*/

type tireNode struct {
	children [26]*tireNode
	value    int
}
type mapSum struct {
	root *tireNode
}

func con() (m mapSum) {
	m.root = new(tireNode)
	return m
}
func (m *mapSum) insert(key string, val int) {
	node := m.root
	chs := []byte(key)
	for i := 0; i < len(chs); i++ {
		ch := chs[i]
		if node.children[ch-'a'] == nil {
			node.children[ch-'a'] = new(tireNode)
		}
		node = node.children[ch-'a']
	}
	node.value = val
}
func (m *mapSum) sum(prefix string) int {
	node := m.root
	chs := []byte(prefix)
	for i := 0; i < len(chs); i++ {
		ch := chs[i]
		if node.children[ch-'a'] == nil {
			return 0
		}
		node = node.children[ch-'a']
	}
	return getsum(node)
}
func getsum(node *tireNode) int {
	if node == nil {
		return 0
	}
	result := node.value
	for _, child := range node.children {
		result += getsum(child)
	}
	return result
}

func main() {
	m := con()
	m.insert("happy", 3)
	fmt.Println(m.sum("hap"))
	m.insert("happen", 2)
	fmt.Println(m.sum("hap"))
}
