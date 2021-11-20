package main

import (
	"fmt"
	"sort"
)

/**
数组中和为0的3个数字
*/
/**
题目：输入一个数组，如何找出数组中所有和为0的3个数字的三元组？需要注意的是，返回值中不得包含重复的三元组。例如，在数组[-1，0，1，2，-1，-4]中有两个三元组的和为0，它们分别是[-1，0，1]和[-1，-1，2]。
*/

/**
操作符 & 取变量地址，使用 * 的意思是以指针的方式间接访问目标对象，打了*的就好比变量的替身，相当于快捷方式
*/
func twoSum1(nums []int, i int, result [][]int) [][]int {
	var result2 [][]int
	j := i + 1
	k := len(nums) - 1
	for {
		if j < k {
			if nums[i]+nums[j]+nums[k] == 0 {
				result2 = append(result2, []int{nums[i], nums[j], nums[k]})
				// 去重
				temp := nums[j]
				for {
					if temp == nums[j] && j < k {
						j++
					} else {
						break
					}
				}
			} else if nums[i]+nums[j]+nums[k] < 0 {
				j++
			} else {
				k--
			}

		} else {
			break
		}
	}
	result = append(result2, result...)
	return result
}

func twoSum2(nums []int, i int, result *[][]int) {
	j := i + 1
	k := len(nums) - 1
	for {
		if j < k {
			if nums[i]+nums[j]+nums[k] == 0 {
				*result = append(*result, []int{nums[i], nums[j], nums[k]})
				// 去重
				temp := nums[j]
				for {
					if temp == nums[j] && j < k {
						j++
					} else {
						break
					}
				}
			} else if nums[i]+nums[j]+nums[k] < 0 {
				j++
			} else {
				k--
			}

		} else {
			break
		}
	}
}
func threeSum(nums []int) [][]int {
	var result [][]int
	if len(nums) >= 3 {
		sort.Ints(nums)
		i := 0
		for {
			if i < len(nums)-2 {
				twoSum2(nums, i, &result)
				temp := nums[i]
				for {
					if i < len(nums) && nums[i] == temp {
						i++
					} else {
						break
					}
				}
			} else {
				break
			}
		}
	}

	return result
}
func main() {
	fmt.Print(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}
