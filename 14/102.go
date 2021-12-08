package main

import "fmt"

/**
加减的目标值
*/
/**
题目：给定一个非空的正整数数组和一个目标值S
，如果为每个数字添加“+”或“-”运算符，请计算有多少种方法可以使这些整数的计算结果为S
。例如，如果输入数组[2，2，2]并且S
等于2，有3种添加“+”或“-”的方法使结果为2，它们分别是2+2-2=2、2-2+2=2及-2+2+2=2。
*/

func findTargetSumWays(nums []int, s int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if (sum+s)%2 == 1 || sum < s {
		return 0
	}
	return subsetSum3(nums, (sum+s)/2)
}
func subsetSum3(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	for _, num := range nums {
		for j := target; j >= num; j-- {
			dp[j] += dp[j-num]
		}
	}
	return dp[target]
}
func main() {
	r := findTargetSumWays([]int{2, 2, 2}, 2)
	fmt.Println(r)
}
