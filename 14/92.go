package main

import (
	"fmt"
	"math"
)

/**
翻转字符
*/
/**
题目：输入一个只包含'0'和'1'的字符串，其中，'0'可以翻转成'1'，'1'可以翻转成'0'。请问至少需要翻转几个字符，才可以使翻转之后的字符串中所有的'0'位于'1'的前面？翻转之后的字符串可能只包含字符'0'或'1'。例如，输入字符串"00110"，至少需要翻转一个字符才能使所有的'0'位于'1'的前面。可以将最后一个字符'0'翻转成'1'，得到字符串"00111"。
*/

func minFlipsMonoIncr(s string) int {
	l := len(s)
	if l == 0 {
		return 0
	}
	dp := [2][2]int{}
	ch := s[0:1]
	if ch == "0" {
		dp[0][0] = 0
		dp[1][0] = 1
	} else {
		dp[0][0] = 1
		dp[1][0] = 0
	}
	for i := 1; i < l; i++ {
		ch = s[i : i+1]
		prev0 := dp[0][(i-1)%2]
		prev1 := dp[1][(i-1)%2]
		if ch == "0" {
			dp[0][i%2] = prev0 + 0
			dp[1][i%2] = int(math.Min(float64(prev0), float64(prev1))) + 1
		} else {
			dp[0][i%2] = prev0 + 1
			dp[1][i%2] = int(math.Min(float64(prev0), float64(prev1))) + 0
		}

	}
	return int(math.Min(float64(dp[0][(l-1)%2]), float64(dp[1][(l-1)%2])))
}
func main() {
	m := minFlipsMonoIncr("00110")
	fmt.Println(m)
}
