package main

import "fmt"

/**
所有子集
 */
/**
题目：输入一个不含重复数字的数据集合，请找出它的所有子集。例如，数据集合[1，2]有4个子集，分别是[]、[1]、[2]和[1，2]。
 */

func subsets(nums []int) [][]int {
	result:=new([][]int)
	if len(nums)==0{
		return *result
	}
	var subset []int
	helper(nums,0,subset,result)
	return *result
}
func helper(nums []int,index int,subset []int,result *[][]int)  {
	if index==len(nums){
		*result=append(*result,subset)
	}else if index<len(nums){
		helper(nums,index+1,subset,result)
		subset=append(subset,nums[index])
		helper(nums,index+1,subset,result)
		subset=subset[:len(subset)-1]
	}
}
func main() {
	r:=subsets([]int{1,2})
	fmt.Println(r)
}

