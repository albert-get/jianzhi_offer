package qian

import (
	"strings"
)

/**

 */

type TrieNode struct {
	Children [26]*TrieNode
	IsWord   bool
}
type Trie struct {
	root *TrieNode
}

func (t *Trie) Insert(word string) {
	node := t.root
	for _, v := range word {
		if node.Children[v-'a'] == nil {
			node.Children[v-'a'] = new(TrieNode)
		}
		node = node.Children[v-'a']
	}
	node.IsWord = true
}
func (t *Trie) Search(word string) bool {
	node := t.root
	for _, v := range word {
		if node.Children[v-'a'] == nil {
			return false
		}
		node = node.Children[v-'a']
	}
	return node.IsWord
}
func (t *Trie) StartsWith(prefix string) bool {
	node := t.root
	for _, v := range prefix {
		if node.Children[v-'a'] == nil {
			return false
		}
		node = node.Children[v-'a']
	}
	return true
}

func BuildTrie(dict []string) *TrieNode {
	root:=new(TrieNode)
	for _,word:=range dict{
		node:=root
		for _,ch:=range word{
			if node.Children[ch-'a']==nil{
				node.Children[ch-'a']=new(TrieNode)
			}
			node=node.Children[ch-'a']
		}
		node.IsWord=true
	}
	return root
}
func BuildTrieR(dict []string) *TrieNode {
	root:=new(TrieNode)
	for _,word:=range dict{
		node:=root
		ch:=[]byte(word)
		for i:=len(ch)-1;i>=0;i--{
			f:=ch[i]
			if node.Children[f-'a']==nil{
				node.Children[f-'a']=new(TrieNode)
			}
			node=node.Children[f-'a']
		}
		node.IsWord=true
	}
	return root
}
func FindPrefix(root *TrieNode,word string) string {
	node:=root
	builder:=""
	for _,ch:=range word{
		if node.IsWord||node.Children[ch-'a']==nil{
			break
		}
		builder+=string(ch)
		node=node.Children[ch-'a']
	}
	if node.IsWord{
		return builder
	}else {
		return ""
	}
}
func ReplaceWords(dict []string,sentence string) string {
	root :=BuildTrie(dict)
	words:=strings.Split(sentence," ")
	for i,ch:=range words{
		prefix:=FindPrefix(root, ch)
		if prefix!=""{
			words[i]=prefix
		}
	}
	return strings.Join(words," ")
}

