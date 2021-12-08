package main

import (
	"fmt"
	"math"
)

/**
子序列的数目
*/
/**
题目：输入字符串S和T，请计算字符串S中有多少个子序列等于字符串T。例如，在字符串"appplep"中，有3个子序列等于字符串"apple"，
*/

func numDistinct(s string, t string) int {
	dp := make([][]int, len(s)+1)
	for i, _ := range dp {
		dp[i] = make([]int, len(t)+1)
	}
	dp[0][0] = 1
	for i := 0; i < len(s); i++ {
		dp[i+1][0] = 1
		for j := 0; j <= i && j < len(t); j++ {
			if s[i] == t[j] {
				dp[i+1][j+1] = dp[i][j] + dp[i][j+1]
			} else {
				dp[i+1][j+1] = dp[i][j+1]
			}
		}
	}
	return dp[len(s)][len(t)]
}
func numDistinct1(s string, t string) int {
	dp := make([]int, len(t)+1)
	if len(s) > 0 {
		dp[0] = 1
	}
	for i := 0; i < len(s); i++ {
		for j := int(math.Min(float64(i), float64(len(t)-1))); j >= 0; j-- {
			if s[i] == t[j] {
				dp[j+1] += dp[j]
			}
		}
	}
	return dp[len(t)]
}
func main() {
	//r := numDistinct("appplep", "apple")
	r := numDistinct1("appplep", "apple")
	fmt.Println(r)
}
