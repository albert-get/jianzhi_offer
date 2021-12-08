package main

import (
	"fmt"
	"math"
)

/**
最少回文分割
*/
/**
题目：输入一个字符串，请问至少需要分割几次才可以使分割出的每个子字符串都是回文？例如，输入字符串"aaba"，至少需要分割1次，从两个相邻字符'a'中间切一刀将字符串分割成两个回文子字符串"a"和"aba"。
*/

func minCut(s string) int {
	l := len(s)
	isPal := make([][]bool, l)
	for i, _ := range isPal {
		isPal[i] = make([]bool, l)
	}
	for i := 0; i < l; i++ {
		for j := 0; j <= i; j++ {
			ch1 := s[i : i+1]
			ch2 := s[j : j+1]
			if ch1 == ch2 && (i <= j+1 || isPal[j+1][i-1]) {
				isPal[j][i] = true
			}
		}
	}
	dp := make([]int, l)
	for i := 0; i < l; i++ {
		if isPal[0][i] {
			dp[i] = 0
		} else {
			dp[i] = i
			for j := 1; j <= i; j++ {
				if isPal[j][i] {
					dp[i] = int(math.Min(float64(dp[i]), float64(dp[i-1]+1)))
				}
			}
		}
	}
	return dp[l-1]
}
func main() {
	r := minCut("aaba")
	fmt.Println(r)
}
