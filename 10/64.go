package main

import (
	"fmt"
	"harry.com/jianzhi_offer/qian"
)

/**
神奇的字典
 */
/**
题目：请实现有如下两个操作的神奇字典。
·函数buildDict，输入单词数组用来创建一个字典。
·函数search，输入一个单词，判断能否修改该单词中的一个字符，使修改之后的单词是字典中的一个单词。
 */

type MagicDictionary struct {
	root *qian.TrieNode
}

func buildDict (dict []string) (d MagicDictionary) {
	d.root=qian.BuildTrie(dict)
	return d
}
func (d *MagicDictionary) search(word string) bool {
	return dfs(d.root,word,0,0)
}
func dfs(root *qian.TrieNode,word string,i int,edit int) bool {
	if root==nil{
		return false
	}
	if root.IsWord&&i==len(word)&&edit==1{
		return true
	}
	c:=[]byte(word)
	if i<len(word)&&edit<=1{
		found:=false
		for j:=0;j<26&&!found;j++{
			var next int
			if j==int(rune(c[i])-'a'){
				next=edit
			}else {
				next=edit+1
			}
			found=dfs(root.Children[j],word,i+1,next)
		}
		return found
	}
	return false
}

func main() {
	d:=buildDict([]string{"happy","new","year"})
	f:=d.search("now")
	fmt.Println(f)
}

