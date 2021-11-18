package main

import "math"

/**
整数除法
*/
/**
题目：输入2个int型整数，它们进行除法计算并返回商，要求不得使用乘号'*'、除号'/'及求余符号'%'。当发生溢出时，返回最大的整数值。假设除数不为0。例如，输入15和2，输出15/2的结果，即7。
 */

func divideCore(dividend int, divisor int) int {
	result := 0
	for {
		if dividend <= divisor {
			value := divisor
			quotient := 1
			for {
				if value >= 0xc0000000 && dividend <= value+value {
					quotient += quotient
					value += value
				} else {
					break
				}
			}
			result += quotient
			dividend -= value
		}else {
			break
		}
	}
	return result
}

func divide(dividend int, divisor int) int {
	if dividend == 0x80000000 && divisor == -1 {
		return math.MaxInt
	}
	negative := 2
	if dividend > 0 {
		negative--
		dividend = -dividend
	}
	if divisor > 0 {
		negative--
		divisor = -divisor
	}
	result := divideCore(dividend, divisor)
	if negative == 1 {
		return -result
	} else {
		return result
	}
}
func main() {
	result := divide(15, 2)
	print(result)
}
