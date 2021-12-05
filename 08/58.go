package main

import (
	"fmt"
	"harry.com/jianzhi_offer/phs"
)

/**
日程表
*/
/**
题目：请实现一个类型MyCalendar用来记录自己的日程安排，该类型用方法book（int start，int end）在日程表中添加一个时间区域为[start，end）的事项（这是一个半开半闭区间）。如果[start，end）中之前没有安排其他事项，则成功添加该事项并返回true；否则，不能添加该事项，并返回false。
*/

type MyCalendar struct {
	events *phs.RBRootMap
}

func ConMyCalendar() MyCalendar {
	return MyCalendar{events: new(phs.RBRootMap)}
}
func (cl *MyCalendar) book(start int, end int) bool {
	event := cl.events.Floor(start)
	if event != nil && event.GetValue() > start {
		return false
	}
	event = cl.events.Ceil(start)
	if event != nil && event.GetValue() < end {
		return false
	}
	cl.events.Set(start, end)
	return true
}
func main() {
	mc := ConMyCalendar()
	fmt.Println(mc.book(10, 20))
	fmt.Println(mc.book(15, 25))
	fmt.Println(mc.book(20, 30))
}
