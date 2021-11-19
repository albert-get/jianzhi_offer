package main

import "fmt"

/**
只出现一次的数字
*/

/**
题目：输入一个整数数组，数组中只有一个数字出现了一次，而其他数字都出现了3次。请找出那个只出现一次的数字。例如，如果输入的数组为[0，1，0，1，0，1，100]，则只出现一次的数字是100。
*/

/**
讲一些计算转化为二进制计算，往往在计算机上有奇效
优化很多算法应该找找数学上的规律
*/

func singleNumber(nums []int) int {
	var bitSums [32]int
	for _, num := range nums {
		for i := 0; i < 32; i++ {
			bitSums[i] += (num >> (31 - i)) & 1 //右移还要与1才能知道第i位
		}
	}
	//bitSums的位置是和数字一样的，右边是低位，左边是高位
	result := 0
	for i := 0; i < 32; i++ {
		result = (result << 1) + bitSums[i]%3
	}
	return result
}

/**
题目：
输入一个整数数组，数组中只有一个数字出现m
次，其他数字都出现n
次。请找出那个唯一出现m
次的数字。假设m
不能被n
整除。
*/
/**
m
不能被n
整除
*/
func singleNumber2(nums []int, n int, m int) int {
	var bitSums [32]int
	for _, num := range nums {
		for i := 0; i < 32; i++ {
			bitSums[i] += (num >> (31 - i)) & 1 //右移还要与1才能知道第i位
		}
	}
	//bitSums的位置是和数字一样的，右边是低位，左边是高位
	result := 0
	for i := 0; i < 32; i++ {
		result = (result << 1) + bitSums[i]%n
	}
	return result / m
}

func main() {
	//fmt.Print(singleNumber([]int{0, 1, 0, 1, 0, 1, 100}))
	fmt.Print(singleNumber2([]int{0, 1, 0, 1, 0, 1, 100, 100}, 3, 2))
}
