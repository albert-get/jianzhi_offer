package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
按权重生成随机数
*/
/**
题目：输入一个正整数数组w
，数组中的每个数字w
[i
]表示下标i
的权重，请实现一个函数pickIndex根据权重比例随机选择一个下标。例如，如果权重数组w
为[1，2，3，4]，那么函数pickIndex将有10%的概率选择0、20%的概率选择1、30%的概率选择2、40%的概率选择3。
*/

type Solution struct {
	sums  []int
	total int
}

func conS(w []int) (s Solution) {
	s.sums = make([]int, len(w))
	for i, item := range w {
		s.total += item
		s.sums[i] = s.total
	}
	return s
}
func (s *Solution) pickIndex() int {
	rand.Seed(time.Now().Unix())
	p := rand.Intn(s.total)
	fmt.Println(p)
	left := 0
	right := len(s.sums)
	for left <= right {
		mid := (left + right) / 2
		if s.sums[mid] > p {
			if mid == 0 || s.sums[mid-1] <= p {
				return mid
			}
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
func main() {
	s := conS([]int{1, 2, 3, 4})
	fmt.Println(s.pickIndex())
}
