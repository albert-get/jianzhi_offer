package main

import (
	"fmt"
	"math"
)

/**
不含重复字符的最长子字符串
 */
/**
题目：输入一个字符串，求该字符串中不含重复字符的最长子字符串的长度。例如，输入字符串"babcca"，其最长的不含重复字符的子字符串是"abc"，长度为3。
 */
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	count := [256]int{}
	i:=0
	j:=-1
	longest:=1
	for ;i<len(s);i++{
		count[s[i]]++
		for hasGreaterThan1(count) {
			j++
			count[s[j]]--
		}
		longest = int(math.Max(float64(longest), float64(i-j)))
	}
	return longest
}
func hasGreaterThan1(counts [256]int) bool {
	for _,v:=range counts{
		if v>1{
			return true
		}
	}
	return false
}
func lengthOfLongestSubstring2(s string) int {
	if len(s)==0 {
		return 0
	}
	counts := [256]int{}
	i:=0
	j:=-1
	longest:=1
	countDup := 0
	for ;i<len(s);i++{
		counts[s[i]]++
		if counts[s[i]]==2{
			countDup++
		}
		for countDup>0{
			j++
			counts[s[j]]--
			if counts[s[i]] == 1{
				countDup --
			}
			longest= int(math.Max(float64(longest), float64(i-j)))
		}
	}
	return longest
}
func main() {
	fmt.Println(lengthOfLongestSubstring("babcca"))
	fmt.Println(lengthOfLongestSubstring2("babcca"))
}
