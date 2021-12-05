package main

import (
	"fmt"
	"math"
)

/**
最大的异或
*/
/**
题目：输入一个整数数组（每个数字都大于或等于0），请计算其中任意两个数字的异或的最大值。例如，在数组[1，3，4，7]中，3和4的异或结果最大，异或结果为7。
*/

type tirenode struct {
	children [2]*tirenode
}

func buildTrie(nums []int32) *tirenode {
	root := new(tirenode)
	for _, num := range nums {
		node := root
		for i := int32(30); i >= 0; i-- {
			bit := (num >> i) & i
			if node.children[bit] == nil {
				node.children[bit] = new(tirenode)
			}
			node = node.children[bit]
		}
	}
	return root
}
func findMaximumXOR(nums []int32) int {
	root := buildTrie(nums)
	max := 0
	for _, num := range nums {
		node := root
		xor := 0

		for i := int32(30); i >= 0; i-- {
			bit := (num >> i) & i
			if node.children[1-bit] != nil {
				xor = xor<<1 + 1
				node = node.children[1-bit]
			} else {
				xor = xor << 1
				node = node.children[bit]
			}
		}
		max = int(math.Max(float64(max), float64(xor)))
	}
	return max
}

func main() {
	max := findMaximumXOR([]int32{1, 3, 4, 7})
	fmt.Println(max)
}
