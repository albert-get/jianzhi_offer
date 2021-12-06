package main

import "fmt"

/**
数组相对排序
*/
/**
题目：输入两个数组arr1和arr2，其中数组arr2中的每个数字都唯一，并且都是数组arr1中的数字。请将数组arr1中的数字按照数组arr2中的数字的相对顺序排序。如果数组arr1中的数字在数组arr2中没有出现，那么将这些数字按递增的顺序排在后面。假设数组中的所有数字都在0到1000的范围内。例如，输入的数组arr1和arr2分别是[2，3，3，7，3，9，2，1，7，2]和[3，2，1]，则数组arr1排序之后为[3，3，3，2，2，2，1，7，7，9]。
*/

func relativeSortArray(arr1 []int, arr2 []int) []int {
	counts := [1001]int{}
	for _, num := range arr1 {
		counts[num]++
	}
	i := 0
	for _, item := range arr2 {
		for counts[item] > 0 {
			arr1[i] = item
			counts[item]--
			i++
		}
	}
	for num := 0; num < len(counts); num++ {
		for counts[num] > 0 {
			arr1[i] = num
			counts[num]--
			i++
		}
	}
	return arr1
}
func main() {
	r := relativeSortArray([]int{2, 3, 3, 7, 3, 9, 2, 1, 7, 2}, []int{3, 2, 1})
	fmt.Println(r)
}
