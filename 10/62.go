package main

import "fmt"

/**
实现前缀树
*/
/**
题目：请设计实现一棵前缀树Trie，它有如下操作。
·函数insert，在前缀树中添加一个字符串。
·函数search，查找字符串。如果前缀树中包含该字符串，则返回true；否则返回false。
·函数startWith，查找字符串前缀。如果前缀树中包含以该前缀开头的字符串，则返回true；否则返回false。
*/

type TrieNode struct {
	children [26]*TrieNode
	isWord   bool
}
type Trie struct {
	root *TrieNode
}

func (t *Trie) insert(word string) {
	node := t.root
	for _, v := range word {
		if node.children[v-'a'] == nil {
			node.children[v-'a'] = new(TrieNode)
		}
		node = node.children[v-'a']
	}
	node.isWord = true
}
func (t *Trie) search(word string) bool {
	node := t.root
	for _, v := range word {
		if node.children[v-'a'] == nil {
			return false
		}
		node = node.children[v-'a']
	}
	return node.isWord
}
func (t *Trie) startsWith(prefix string) bool {
	node := t.root
	for _, v := range prefix {
		if node.children[v-'a'] == nil {
			return false
		}
		node = node.children[v-'a']
	}
	return true
}

func main() {
	t := Trie{root: &TrieNode{}}
	t.insert("boy")
	t.insert("boss")
	t.insert("cowboy")
	t.insert("cow")
	fmt.Println(t)
}
