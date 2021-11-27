package main

import (
	"fmt"
	"math"
)

/**
0和1个数相同的子数组
*/
/**
题目：输入一个只包含0和1的数组，请问如何求0和1的个数相同的最长连续子数组的长度？例如，在数组[0，1，0]中有两个子数组包含相同个数的0和1，分别是[0，1]和[1，0]，它们的长度都是2，因此输出2。
*/

func findMaxLength(nums []int) int {
	sumToIndex := make(map[int]int)
	sumToIndex[0] = -1
	sum := 0
	maxLength := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			sum += -1
		} else {
			sum += nums[i]
		}
		index, ok := sumToIndex[sum]
		if ok {
			maxLength = int(math.Max(float64(i-index), float64(maxLength)))
		} else {
			sumToIndex[sum] = i
		}
	}
	return maxLength
}
func main() {
	fmt.Print(findMaxLength([]int{0, 1, 0}))
}
