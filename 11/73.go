package main

import (
	"fmt"
	"math"
)

/**
狒狒吃香蕉
*/
/**
题目：狒狒很喜欢吃香蕉。一天它发现了n
堆香蕉，第i
堆有piles[i
]根香蕉。门卫刚好走开，H
小时后才会回来。狒狒吃香蕉喜欢细嚼慢咽，但又想在门卫回来之前吃完所有的香蕉。请问狒狒每小时至少吃多少根香蕉？如果狒狒决定每小时吃k
根香蕉，而它在吃的某一堆剩余的香蕉的数目少于k
，那么它只会将这一堆的香蕉吃完，下一个小时才会开始吃另一堆的香蕉。
*/

func minEatingSpeed(piles []int, H int) int {
	max := math.MinInt
	for _, item := range piles {
		max = int(math.Max(float64(max), float64(item)))
	}
	left := 1
	right := max
	for left <= right {
		mid := (left + right) / 2
		hours := getHours(piles, mid)
		if hours <= H {
			if mid == 1 || getHours(piles, mid-1) > H {
				return mid
			}
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
func getHours(piles []int, speed int) int {
	hours := 0
	for _, pile := range piles {
		hours += (pile + speed - 1) / speed
	}
	return hours
}
func main() {
	h := minEatingSpeed([]int{3, 6, 7, 11}, 8)
	fmt.Println(h)
}
