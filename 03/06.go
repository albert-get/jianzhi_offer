package main

import "fmt"

/**
最多删除一个字符得到回文
*/
/**
题目：给定一个字符串，请判断如果最多从字符串中删除一个字符能不能得到一个回文字符串。例如，如果输入字符串"abca"，由于删除字符'b'或'c'就能得到一个回文字符串，因此输出为true。
*/
func validPalindrome(s string) bool {
	start := 0
	end := len(s) - 1
	for start < len(s)/2 {
		if s[start] != s[end] {
			break
		}
		start++
		end--
	}
	return start == len(s)/2 || isPalindrome2(s, start, end-1) || isPalindrome2(s, start+1, end)
}
func isPalindrome2(s string, start int, end int) bool {
	for start < end {
		if s[start] != s[end] {
			break
		}
		start++
		end--
	}
	return start >= end
}
func main() {
	fmt.Println(validPalindrome("abca"))
}
