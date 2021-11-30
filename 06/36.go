package main

import (
	"fmt"
	"harry.com/jianzhi_offer/zhan"
	"strconv"
)

/**
后缀表达式
*/
/**
题目：后缀表达式是一种算术表达式，它的操作符在操作数的后面。输入一个用字符串数组表示的后缀表达式，请输出该后缀表达式的计算结果。假设输入的一定是有效的后缀表达式。例如，后缀表达式["2"，"1"，"3"，"*"，"+"]对应的算术表达式是“2+1*3”，因此输出它的计算结果5。
*/

func evalRPN(tokens []string) int {
	stack := zhan.NewStack()
	for _, token := range tokens {

		switch token {
		case "+", "-", "*", "/":
			num1, _ := stack.Pop().(string)
			num2, _ := stack.Pop().(string)
			n1, _ := strconv.Atoi(num1)
			n2, _ := strconv.Atoi(num2)
			stack.Push(strconv.Itoa(calculate(n1, n2, token)))
		default:
			stack.Push(token)
		}
	}
	res, _ := stack.Pop().(string)
	r, _ := strconv.Atoi(res)
	return r
}
func calculate(num1 int, num2 int, operator string) int {
	switch operator {
	case "+":
		return num1 + num2
	case "-":
		return num1 - num2
	case "*":
		return num1 * num2
	case "/":
		return num1 / num2
	default:
		return 0
	}
}
func main() {
	fmt.Println(evalRPN([]string{"2", "1", "3", "*", "+"}))
}
