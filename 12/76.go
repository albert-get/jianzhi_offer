package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
数组中第k大的数字
*/
/**
题目：请从一个乱序数组中找出第k
大的数字。例如，数组[3，1，2，4，5，5，6]中第3大的数字是5。
*/

func sortArray(nums []int) []int {
	quicksort(nums, 0, len(nums)-1)
	return nums
}
func quicksort(nums []int, start int, end int) {
	if end < start {
		pivot := partition(nums, start, end)
		quicksort(nums, start, pivot-1)
		quicksort(nums, pivot+1, end)
	}
}
func partition(nums []int, start int, end int) int {
	rand.Seed(time.Now().Unix())
	random := rand.Intn(end-start+1) + start
	swap(nums, random, end)
	small := start - 1
	for i := start; i < end; i++ {
		if nums[i] < nums[end] {
			small++
			swap(nums, i, small)
		}
	}
	small++
	swap(nums, small, end)
	return small
}
func swap(nums []int, index1 int, index2 int) {
	if index1 != index2 {
		nums[index1], nums[index2] = nums[index2], nums[index1]
	}
}
func findKthLargest(nums []int, k int) int {
	target := len(nums) - k
	start := 0
	end := len(nums) - 1
	index := partition(nums, start, end)
	for index != target {
		if index > target {
			end = index - 1
		} else {
			start = index + 1
		}
		index = partition(nums, start, end)
	}
	return nums[index]
}
func main() {
	k := findKthLargest([]int{3, 1, 2, 4, 5, 5, 6}, 3)
	fmt.Println(k)
}
