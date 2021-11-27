package main

import "fmt"

/**
回文子字符串的个数
*/
/**
题目：给定一个字符串，请问该字符串中有多少个回文连续子字符串？例如，字符串"abc"有3个回文子字符串，分别为"a"、"b"和"c"；而字符串"aaa"有6个回文子字符串，分别为"a"、"a"、"a"、"aa"、"aa"和"aaa"。
*/
func countSubstrings(s string) int {
	if len(s) == 0 {
		return 0
	}
	count := 0
	for i := 0; i < len(s); i++ {
		count += countPalindrome(s, i, i)
		count += countPalindrome(s, i, i+1)
	}
	return count
}
func countPalindrome(s string, start int, end int) int {
	count := 0
	for start >= 0 && end < len(s) && s[start] == s[end] {
		count++
		start--
		end++
	}
	return count
}
func main() {
	fmt.Println(countSubstrings("aaa"))
}
