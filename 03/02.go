package main

import "fmt"

/**
字符串中的所有变位词
*/
/**
题目：输入字符串s1和s2，如何找出字符串s2的所有变位词在字符串s1中的起始下标？假设两个字符串中只包含英文小写字母。例如，字符串s1为"cbadabacg"，字符串s2为"abc"，字符串s2的两个变位词"cba"和"bac"是字符串s1中的子字符串，输出它们在字符串s1中的起始下标0和5。
*/

func findAnagrams(c string, sp string) []int {
	var indices []int
	if len(sp) < len(c) {
		return indices
	}
	counts := [26]int{}
	for i, v := range c {
		counts[v-'a']++
		counts[sp[i]-'a']--
	}
	if areAllZero1(counts) {
		indices = append(indices, 0)
	}
	for i := len(c); i < len(sp); i++ {
		counts[sp[i]-'a']--
		counts[sp[i-len(c)]-'a']++
		if areAllZero1(counts) {
			indices = append(indices, i-len(c)+1)
		}
	}
	return indices

}

func areAllZero1(counts [26]int) bool {
	for _, v := range counts {
		if v != 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(findAnagrams("abc", "cbadabacg"))
}
