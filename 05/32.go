package main

import "fmt"

/**
有效的变位词
*/
/**
题目：给定两个字符串s和t，请判断它们是不是一组变位词。在一组变位词中，它们中的字符及每个字符出现的次数都相同，但字符的顺序不能相同。例如，"anagram"和"nagaram"就是一组变位词。
*/

func isAnagram(str1 string, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}
	counts := [26]int{}
	for _, v := range str1 {
		counts[v-'a']++
	}
	for _, k := range str2 {
		if counts[k-'a'] == 0 {
			return false
		}
		counts[k-'a']--
	}
	return true
}
func isAnagram2(str1 string, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}
	counts := make(map[rune]int)
	for _, v := range str1 {
		counts[v]++
	}
	for _, k := range str2 {
		n, ok := counts[k]
		if !ok || n == 0 {
			return false
		}
		counts[k]--
	}
	return true
}
func main() {
	fmt.Println(isAnagram("anagram", "nagaram"))
	fmt.Println(isAnagram2("辣不怕", "不怕辣"))
}
