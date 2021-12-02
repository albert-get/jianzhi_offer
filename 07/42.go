package main

import (
	"fmt"
	"harry.com/jianzhi_offer/dui"
)

/**
最近请求次数
*/
/**
题目：请实现如下类型RecentCounter，它是统计过去3000ms内的请求次数的计数器。该类型的构造函数RecentCounter初始化计数器，请求数初始化为0；函数ping（int t）在时间t添加一个新请求（t表示以毫秒为单位的时间），并返回过去3000ms内（时间范围为[t-3000，t]）发生的所有请求数。假设每次调用函数ping的参数t都比之前调用的参数值大。
*/

type RecentCounter struct {
	times      dui.QueueInt
	windowSize int
}

func Constructor1() *RecentCounter {
	return &RecentCounter{
		times:      dui.QueueInt{},
		windowSize: 3000,
	}
}
func (r *RecentCounter) ping(t int) int {
	r.times.Push(t)
	for r.times.Peek()+r.windowSize < t {
		r.times.Pop()
	}
	return r.times.Len()
}
func main() {
	rc := Constructor1()
	fmt.Println(rc.ping(1))
	fmt.Println(rc.ping(10))
	fmt.Println(rc.ping(3001))
	fmt.Println(rc.ping(3002))
}
