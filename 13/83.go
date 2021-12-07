package main

import "fmt"

/**
没有重复元素集合的全排列
 */
/**
题目：给定一个没有重复数字的集合，请找出它的所有全排列。例如，集合[1，2，3]有6个全排列，分别是[1，2，3]、[1，3，2]、[2，1，3]、[2，3，1]、[3，1，2]和[3，2，1]。
 */

func permute(nums []int) [][]int {
	result:=new([][]int)
	helper3(nums,0,result)
	return *result
}
func helper3(nums []int,i int,result *[][]int)  {
	if i==len(nums){
		var permutation []int
		for _,num:=range nums{
			permutation=append(permutation,num)
		}
		*result=append(*result,permutation)
	}else {
		for j:=i;j<len(nums);j++{
			nums[i],nums[j]=nums[j],nums[i]
			helper3(nums,i+1,result)
			nums[i],nums[j]=nums[j],nums[i]
		}
	}
}
func main() {
	r:=permute([]int{1,2,3})
	fmt.Println(r)
}
