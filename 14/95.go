package main

import (
	"fmt"
	"math"
)

/**
最长公共子序列
*/
/**
题目：输入两个字符串，请求出它们的最长公共子序列的长度。如果从字符串s1中删除若干字符之后能得到字符串s2，那么字符串s2就是字符串s1的一个子序列。例如，从字符串"abcde"中删除两个字符之后能得到字符串"ace"，因此字符串"ace"是字符串"abcde"的一个子序列。但字符串"aec"不是字符串"abcde"的子序列。如果输入字符串"abcde"和"badfe"，那么它们的最长公共子序列是"bde"，因此输出3。
*/

func longestCommonSubsequence(text1 string, text2 string) int {
	len1 := len(text1)
	len2 := len(text2)
	dp := make([][]int, len1+1)
	for i, _ := range dp {
		dp[i] = make([]int, len2+1)
	}
	for i := 0; i < len1; i++ {
		for j := 0; j < len2; j++ {
			if text1[i] == text2[j] {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = int(math.Max(float64(dp[i][j+1]), float64(dp[i+1][j])))
			}
		}
	}
	return dp[len1][len2]
}
func longestCommonSubsequence1(text1 string, text2 string) int {
	len1 := len(text1)
	len2 := len(text2)
	if len1 < len2 {
		return longestCommonSubsequence1(text2, text1)
	}
	dp := [2][]int{}
	for i, _ := range dp {
		dp[i] = make([]int, len2+1)
	}
	for i := 0; i < len1; i++ {
		for j := 0; j < len2; j++ {
			if text1[i] == text2[j] {
				dp[(i+1)%2][j+1] = dp[i%2][j] + 1
			} else {
				dp[(i+1)%2][j+1] = int(math.Max(float64(dp[i%2][j+1]), float64(dp[(i+1)%2][j])))
			}
		}
	}
	return dp[len1%2][len2]
}
func longestCommonSubsequence2(text1 string, text2 string) int {
	len1 := len(text1)
	len2 := len(text2)
	if len1 < len2 {
		return longestCommonSubsequence2(text2, text1)
	}
	dp := make([]int, len2+1)
	for i := 0; i < len1; i++ {
		prev := dp[0]
		for j := 0; j < len2; j++ {
			var cur int
			if text1[i] == text2[j] {
				cur = prev + 1
			} else {
				cur = int(math.Max(float64(dp[j]), float64(dp[j+1])))
			}
			prev = dp[j+1]
			dp[j+1] = cur
		}
	}
	return dp[len2]
}
func main() {
	//r:=longestCommonSubsequence("abcde","badfe")
	//r:=longestCommonSubsequence1("abcde","badfe")
	r := longestCommonSubsequence2("abcde", "badfe")
	fmt.Println(r)
}
