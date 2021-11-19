package main

/**
二进制加法
*/
/**
题目：输入两个表示二进制的字符串，请计算它们的和，并以二进制字符串的形式输出。例如，输入的二进制字符串分别是"11"和"10"，则输出"101"。
*/

func addBinary(a string, b string) string {
	var result string
	i := len(a) - 1
	j := len(b) - 1
	var carry uint8 = 0
	for {
		if i >= 0 || j >= 0 {
			var (
				digitA uint8
				digitB uint8
			)
			if i >= 0 {
				digitA = a[i] - '0'
				i = i - 1
			}
			if j >= 0 {
				digitB = b[j] - '0'
				j = j - 1
			}
			sum := digitA + digitB + carry
			var pre string
			if sum >= 2 {
				carry = 1
				pre = "0"
			} else {
				carry = 0
				pre = "1"
			}
			result += pre
		} else {
			break
		}
	}
	if carry == 1 {
		result += "1"
	}
	var sa []rune
	for _, k := range []rune(result) {
		sa = append(sa, k)
	}
	return string(sa)
}

func main() {
	print(addBinary("1011", "10111101"))
}
