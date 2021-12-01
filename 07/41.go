package main

import (
	"fmt"
	"harry.com/jianzhi_offer/dui"
)

/**
滑动窗口的平均值
*/
/**
题目：请实现如下类型MovingAverage，计算滑动窗口中所有数字的平均值，该类型构造函数的参数确定滑动窗口的大小，每次调用成员函数next时都会在滑动窗口中添加一个整数，并返回滑动窗口中所有数字的平均值。
*/

type MovingAverage struct {
	nums     dui.QueueInt
	capacity int
	sum      int
}

func Constructor(size int) *MovingAverage {
	return &MovingAverage{
		capacity: size,
		nums:     dui.QueueInt{},
	}
}
func (mva *MovingAverage) next(val int) float32 {
	mva.nums.Push(val)
	mva.sum += val
	if mva.nums.Len() > mva.capacity {
		mva.sum -= mva.nums.Pop()
	}
	return float32(mva.sum) / float32(mva.nums.Len())
}

func main() {
	q := Constructor(3)

	fmt.Println(q.next(2))
	fmt.Println(q.next(3))
	fmt.Println(q.next(4))
}
