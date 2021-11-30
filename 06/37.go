package main

import (
	"fmt"
	"harry.com/jianzhi_offer/zhan"
)

/**
小行星碰撞
*/
/**
题目：输入一个表示小行星的数组，数组中每个数字的绝对值表示小行星的大小，数字的正负号表示小行星运动的方向，正号表示向右飞行，负号表示向左飞行。如果两颗小行星相撞，那么体积较小的小行星将会爆炸最终消失，体积较大的小行星不受影响。如果相撞的两颗小行星大小相同，那么它们都会爆炸消失。飞行方向相同的小行星永远不会相撞。求最终剩下的小行星。例如，有6颗小行星[4，5，-6，4，8，-5]，如图6.2所示（箭头表示飞行的方向），它们相撞之后最终剩下3颗小行星[-6，4，8]。
*/

func steroidCollision(asteroids []int) []int {
	stack := zhan.NewStackInt()
	for _, v := range asteroids {
		for !stack.IsEmpty() && stack.Peek() > 0 && stack.Peek() < (-v) {
			stack.Pop()
		}
		if !stack.IsEmpty() && v < 0 && stack.Peek() == (-v) {
			stack.Pop()
		} else if v > 0 || stack.IsEmpty() || stack.Peek() < 0 {
			stack.Push(v)
		}
	}
	var res []int
	for !stack.IsEmpty() {
		res = append(res, stack.Pop())
	}
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}
func main() {
	fmt.Println(steroidCollision([]int{4, 5, -6, 4, 8, -5}))
}
