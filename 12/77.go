package main

import (
	"fmt"
	"math"
)

/**
归并
*/

func sortArray1(nums []int) []int {
	length := len(nums)
	src := nums
	dst := make([]int, length)
	for seg := 1; seg < length; seg += seg {
		for start := 0; start < length; start += seg * 2 {
			mid := int(math.Min(float64(start+seg), float64(length)))
			end := int(math.Min(float64(start+seg*2), float64(length)))
			i := start
			j := mid
			k := start
			for i < mid || j < end {
				if j == end || (i < mid && src[i] < src[j]) {
					dst[k] = src[i]
					k++
					i++
				} else {
					dst[k] = src[j]
					k++
					j++
				}
			}
		}
		temp := src
		src = dst
		dst = temp
	}
	return src
}

func sortArray2(nums []int) []int {
	dst := make([]int, len(nums))
	copy(dst, nums)
	mergesort(nums, dst, 0, len(nums))
	return dst
}
func mergesort(src []int, dst []int,start int, end int) {
	if start+1 >= end {
		return
	}
	mid := (start + end) / 2
	mergesort(dst, src, start, mid)
	mergesort(dst, src, mid, end)
	i := start
	j := end
	k := start
	for i < mid || j < end {
		if j == end || (i < mid && src[i] < src[j]) {
			dst[k] = src[i]
			k++
			i++
		} else {
			dst[k] = src[j]
			k++
			j++
		}
	}
}
func main() {
	//r := sortArray1([]int{4, 1, 5, 6, 2, 7, 8, 3})
	r := sortArray2([]int{4, 1, 5, 6, 2, 7, 8, 3})
	fmt.Println(r)
}
