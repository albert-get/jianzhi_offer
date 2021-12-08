package main

import (
	"fmt"
	"math"
)

/**
最长斐波那契数列
*/
/**
题目：输入一个没有重复数字的单调递增的数组，数组中至少有3个数字，请问数组中最长的斐波那契数列的长度是多少？例如，如果输入的数组是[1，2，3，4，5，6，7，8]，由于其中最长的斐波那契数列是1、2、3、5、8，因此输出是5。
*/

func lenLongestFibSubseq(A []int) int {
	mymap := make(map[int]int)
	for i, v := range A {
		mymap[v] = i
	}
	dpc := make([]int, len(A))
	for m := 0; m < len(A); m++ {
		dpc[m] = -1
	}
	dp := make([][]int, len(A))
	for m := 0; m < len(A); m++ {
		t := make([]int, len(A))
		copy(t, dpc)
		dp[m] = t
	}
	result := 2
	for i := 1; i < len(A); i++ {
		for j := 0; j < i; j++ {
			k := mymap[A[i]-A[j]]
			if k >= 0 && k < j {
				dp[i][j] = dp[j][k] + 1
			} else {
				dp[i][j] = 2
			}
			result = int(math.Max(float64(result), float64(dp[i][j])))
		}
	}
	fmt.Println(dp)
	if result > 2 {
		return result
	} else {
		return 0
	}
}
func main() {
	r := lenLongestFibSubseq([]int{1, 2, 3, 4, 5, 6, 7, 8})
	fmt.Println(r)
}
