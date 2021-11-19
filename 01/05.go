package main

import (
	"fmt"
	"math"
)

/**
单词长度的最大乘积
*/

/**
题目：输入一个字符串数组words，请计算不包含相同字符的两个字符串words[i]和words[j]的长度乘积的最大值。
如果所有字符串都包含至少一个相同字符，那么返回0。假设字符串中只包含英文小写字母。
例如，输入的字符串数组words为["abcw"，"foo"，"bar"，"fxyz"，"abcdef"]，数组中的字符串"bar"与"foo"没有相同的字符，它们长度的乘积为9。"abcw"与"fxyz"也没有相同的字符，它们长度的乘积为16，这是该数组不包含相同字符的一对字符串的长度乘积的最大值。
*/

func maxProduct(words []string) int {
	flags := make([][26]bool, len(words))
	for i := 0; i < len(words); i++ {
		for _, c := range words[i] {
			flags[i][c-'a'] = true
		}
	}
	result := 0
	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words); j++ {
			k := 0
			for ; k < 26; k++ {
				if flags[i][k] && flags[j][k] {
					break
				}
			}
			if k == 26 {
				prod := len(words[i]) * len(words[j])
				result = int(math.Max(float64(prod), float64(result)))
			}
		}
	}
	return result
}
func maxProduct2(words []string) int {
	flags := make([]int32, len(words))
	for i := 0; i < len(words); i++ {
		for _, c := range words[i] {
			flags[i] |= 1 << (c - 'a')
		}
	}
	result := 0
	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words); j++ {
			if flags[i]&flags[j] == 0 {
				prod := len(words[i]) * len(words[j])
				result = int(math.Max(float64(prod), float64(result)))
			}
		}
	}
	return result
}

func main() {
	fmt.Print(maxProduct2([]string{"abcw", "foo", "bar", "fxyz", "abcdef"}))
}
