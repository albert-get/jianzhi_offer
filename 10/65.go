package main

import (
	"fmt"
	"harry.com/jianzhi_offer/qian"
)

/**
最短的单词编码
*/
/**
题目：输入一个包含n
个单词的数组，可以把它们编码成一个字符串和n
个下标。例如，单词数组["time"，"me"，"bell"]可以编码成一个字符串"time＃bell＃"，然后这些单词就可以通过下标[0，2，5]得到。对于每个下标，都可以从编码得到的字符串中相应的位置开始扫描，直到遇到'＃'字符前所经过的子字符串为单词数组中的一个单词。例如，从"time＃bell＃"下标为2的位置开始扫描，直到遇到'＃'前经过子字符串"me"是给定单词数组的第2个单词。给定一个单词数组，请问按照上述规则把这些单词编码之后得到的最短字符串的长度是多少？如果输入的是字符串数组["time"，"me"，"bell"]，那么编码之后最短的字符串是"time＃bell＃"，长度是10。
*/

func minimumLengthEncoding(words []string) int {
	root := qian.BuildTrieR(words)
	total := []int{0}
	dfs1(root, 1, total)
	return total[0]
}
func dfs1(root *qian.TrieNode, length int, total []int) {
	isLeaf := true
	for _, child := range root.Children {
		if child != nil {
			isLeaf = false
			dfs1(child, length+1, total)
		}
	}
	if isLeaf {
		total[0] += length
	}
}

func Reverse(s string) string {
	a := func(s string) *[]rune {
		var b []rune
		for _, k := range []rune(s) {
			defer func(v rune) {
				b = append(b, v)
			}(k)
		}
		return &b
	}(s)
	return string(*a)
}

func main() {

	n := minimumLengthEncoding([]string{"at", "bat", "cat"})
	fmt.Println(n)
}
