package main

import (
	"fmt"
	"harry.com/jianzhi_offer/phs"
)

/**
值和下标之差都在给定的范围内
*/
/**
题目：给定一个整数数组nums和两个正数k、t，请判断是否存在两个不同的下标i和j满足i和j之差的绝对值不大于给定的k，并且两个数值nums[i]和nums[j]的差的绝对值不大于给定的t。
*/

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	var root *phs.RBRootSet = new(phs.RBRootSet)
	for i, v := range nums {
		lower := root.Floor(v)
		if lower >= v-t {
			return true
		}
		upper := root.Ceil(v)
		if upper <= v+t {
			return true
		}
		root.Set(v)
		if i >= k {
			root.Delete(nums[i-k])
		}
	}
	return false
}

func containsNearbyAlmostDuplicate2(nums []int, k int, t int) bool {
	buckets := make(map[int]int)
	bucketSize := t + 1
	for i, num := range nums {
		id := getBucketID(num, bucketSize)
		_, idok := buckets[id]
		idj1, idj1ok := buckets[id-1]
		idd1, idd1ok := buckets[id+1]
		if idok || (idj1ok && idj1+t >= num) || (idd1ok && idd1-t <= num) {
			return true
		}
		buckets[id] = num
		if i >= k {
			delete(buckets, getBucketID(nums[i-k], bucketSize))
		}
	}
	return false
}
func getBucketID(num int, bucketSize int) int {
	if num >= 0 {
		return int(num / bucketSize)
	} else {
		return int((num + 1) / (bucketSize - 1))
	}
}

func main() {
	//var datas = []phs.Type{10, 40, 30, 60, 90, 70, 20, 80}
	//var root *phs.RBRoot = new(phs.RBRoot)
	//for _, data := range datas {
	//	var node = &phs.RBNode{Key:data}
	//	root.Add(node)
	//}
	//fmt.Println(root)

	//var datas = []phs.Type{10, 40, 30, 60, 90, 70, 20, 80}
	//var root *phs.RBRootSet = new(phs.RBRootSet)
	//for _, data := range datas {
	//	root.Set(data)
	//}
	//root.Set(50)
	//fmt.Println(root)

	//var datas = [][2]phs.Type{{1,2},{10,30},{2,50}}
	//var root *phs.RBRootMap = new(phs.RBRootMap)
	//for _, data := range datas {
	//	root.Set(data[0],data[1])
	//}
	//fmt.Println(root)

	//flag:=containsNearbyAlmostDuplicate([]int{1,2,3,1},3,0)
	//flag := containsNearbyAlmostDuplicate([]int{1, 5, 9, 1, 5, 9}, 2, 3)
	//flag:=containsNearbyAlmostDuplicate2([]int{1,2,3,1},3,0)
	flag := containsNearbyAlmostDuplicate2([]int{1, 5, 9, 1, 5, 9}, 2, 3)
	fmt.Println(flag)
}
