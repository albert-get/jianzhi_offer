package main

import (
	"fmt"
	"math/rand"
)

/**
插入、删除和随机访问都是O（1）的容器
*/
/**
题目：设计一个数据结构，使如下3个操作的时间复杂度都是O（1）。
·insert（value）：如果数据集中不包含一个数值，则把它添加到数据集中。
·remove（value）：如果数据集中包含一个数值，则把它删除。
·getRandom()：随机返回数据集中的一个数值，要求数据集中每个数字被返回的概率都相同。
*/

type RandomizedSet struct {
	numToLocation map[int]int
	nums          []int
}

func (this *RandomizedSet) insert(val int) bool {
	if _, ok := this.numToLocation[val]; ok {
		return false
	}
	this.numToLocation[val] = len(this.nums)
	this.nums = append(this.nums, val)
	return true
}
func (this *RandomizedSet) remove(val int) bool {
	if _, ok := this.numToLocation[val]; !ok {
		return false
	}
	location := this.numToLocation[val]
	this.numToLocation[this.nums[len(this.nums)-1]] = location
	delete(this.numToLocation, val)
	this.nums[location] = this.nums[len(this.nums)-1]
	this.nums = this.nums[:len(this.nums)-2]
	return true
}
func (this *RandomizedSet) getRandom() int {
	return this.nums[rand.Intn(len(this.nums))]
}
func Contrucor(arr []int) *RandomizedSet {
	mymap := RandomizedSet{nums: arr}
	tempmap := make(map[int]int)
	for i, v := range arr {
		tempmap[v] = i
	}
	mymap.numToLocation = tempmap
	return &mymap
}

func main() {
	mapforme := Contrucor([]int{1, 2, 3})
	fmt.Println(mapforme.remove(3))
	fmt.Println(mapforme)
}
