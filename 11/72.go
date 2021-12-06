package main

import "fmt"

/**
求平方根
*/
/**
题目：输入一个非负整数，请计算它的平方根。正数的平方根有两个，只输出其中的正数平方根。如果平方根不是整数，那么只需要输出它的整数部分。例如，如果输入4则输出2；如果输入18则输出4。
*/

func mySqrt(n int) int {
	left := 1
	right := n
	for left <= right {
		mid := left + (right-left)/2
		if mid <= n/mid {
			if (mid + 1) > n/(mid+1) {
				return mid
			}
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return 0
}
func main() {
	s:=mySqrt(18)
	fmt.Println(s)
}
