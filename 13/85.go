package main

import "fmt"

/**
生成匹配的括号
*/
/**
题目：输入一个正整数n
，请输出所有包含n
个左括号和n
个右括号的组合，要求每个组合的左括号和右括号匹配。例如，当n
等于2时，有两个符合条件的括号组合，分别是"（()）"和"()()"。
*/

func generateParenthesis(n int) []string {
	result := new([]string)
	helper6(n, n, "", result)
	return *result
}
func helper6(left int, right int, parentesis string, result *[]string) {
	if left == 0 && right == 0 {
		*result = append(*result, parentesis)
		return
	}
	if left > 0 {
		helper6(left-1, right, parentesis+"(", result)
	}
	if left < right {
		helper6(left, right-1, parentesis+")", result)
	}
}
func main() {
	r := generateParenthesis(2)
	fmt.Println(r)
}
