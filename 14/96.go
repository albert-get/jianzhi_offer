package main

import "fmt"

/**
字符串交织
*/
/**
题目：输入3个字符串s1、s2和s3，请判断字符串s3能不能由字符串s1和s2交织而成，即字符串s3的所有字符都是字符串s1或s2中的字符，字符串s1和s2中的字符都将出现在字符串s3中且相对位置不变。例如，字符串"aadbbcbcac"可以由字符串"aabcc"和"dbbca"交织而成，
*/

func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	dp := make([][]bool, len(s1)+1)
	for i, _ := range dp {
		dp[i] = make([]bool, len(s2)+1)
	}
	dp[0][0] = true
	for i := 0; i < len(s1); i++ {
		if s1[i] == s3[i] && dp[i][0] {
			dp[i+1][0] = true
		}
	}
	for j := 0; j < len(s2); j++ {
		if s2[j] == s3[j] && dp[0][j] {
			dp[0][j+1] = true
		}
	}
	for i := 0; i < len(s1); i++ {
		for j := 0; j < len(s2); j++ {
			ch1 := s1[i]
			ch2 := s2[j]
			ch3 := s3[i+j+1]
			if (ch1 == ch3 && dp[i][j+1]) || (ch2 == ch3 && dp[i+1][j]) {
				dp[i+1][j+1] = true
			}
		}
	}
	return dp[len(s1)][len(s2)]
}
func isInterleave1(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	if len(s1) < len(s2) {
		return isInterleave1(s2, s1, s3)
	}
	dp := make([]bool, len(s2)+1)
	dp[0] = true
	for j := 0; j < len(s2); j++ {
		if s2[j] == s3[j] && dp[j] {
			dp[j+1] = true
		}
	}
	for i := 0; i < len(s1); i++ {
		if dp[0] && s1[i] == s3[i] {
			dp[0] = true
		}
		for j := 0; j < len(s2); j++ {
			ch1 := s1[i]
			ch2 := s2[j]
			ch3 := s3[i+j+1]
			if (ch1 == ch3 && dp[j+1]) || (ch2 == ch3 && dp[j]) {
				dp[j+1] = true
			}
		}
	}
	return dp[len(s2)]
}
func main() {
	//r:=isInterleave("aabcc","dbbca","aadbbcbcac")
	r := isInterleave1("aabcc", "dbbca", "aadbbcbcac")
	fmt.Println(r)
}
