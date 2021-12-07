package main

import "fmt"

/**
分割回文子字符串
*/
/**
题目：输入一个字符串，要求将它分割成若干子字符串，使每个子字符串都是回文。请列出所有可能的分割方法。例如，输入"google"，将输出3种符合条件的分割方法，分别是["g"，"o"，"o"，"g"，"l"，"e"]、["g"，"oo"，"g"，"l"，"e"]和["goog"，"l"，"e"]。
*/

func partition(s string) [][]string {
	result := new([][]string)
	helper7(s, 0, []string{}, result)
	return *result
}
func helper7(str string, start int, substrings []string, result *[][]string) {
	if start == len(str) {
		*result = append(*result, substrings)
		return
	}
	for i := start; i < len(str); i++ {
		if isPalindrome(str, start, i) {
			substrings = append(substrings, str[start:i+1])
			helper7(str, i+1, substrings, result)
			substrings = substrings[:len(substrings)-1]
		}
	}
}
func isPalindrome(str string, start int, end int) bool {
	for start < end {
		if str[start] != str[end] {
			return false
		}
		start++
		end--
	}
	return true
}
func main() {
	r := partition("google")
	fmt.Println(r)
}
