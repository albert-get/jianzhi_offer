package main

import "fmt"

/**
包含k个元素的组合
*/
/**
题目：输入n
和k
，请输出从1到n
中选取k
个数字组成的所有组合。例如，如果n
等于3，k
等于2，将组成3个组合，分别是[1，2]、[1，3]和[2，3]。
*/

func combine(n int, k int) [][]int {
	result := new([][]int)
	var combination []int
	helper1(n, k, 1, combination, result)
	return *result
}
func helper1(n int, k int, i int, combination []int, result *[][]int) {
	if len(combination) == k {
		*result = append(*result, combination)
	} else if i <= n {
		helper1(n, k, i+1, combination, result)
		combination = append(combination, i)
		helper1(n, k, i+1, combination, result)
		combination = combination[:len(combination)-1]
	}
}

func main() {
	r := combine(3, 2)
	fmt.Println(r)
}
