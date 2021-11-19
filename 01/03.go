package main

/**
前n个数字二进制形式中1的个数
*/

/**
题目：输入一个非负数n，请计算0到n之间每个数字的二进制形式中1的个数，并输出一个数组。例如，输入的n为4，由于0、1、2、3、4的二进制形式中1的个数分别为0、1、1、2、1，因此输出数组[0，1，1，2，1]。
*/

/**

 */
//第一步叫去直觉
func countBits1(num uint) []uint {
	//var result []uint // 1
	result := make([]uint, num+1) // 2
	var i uint = 0
	for ; i <= num; i++ {
		j := i
		//result = append(result,0) // 1
		for {
			if j != 0 {
				result[i]++
				j = j & (j - 1)
			} else {
				break
			}
		}
	}
	return result
}

// 第二步找规律
func countBits2(num uint) []uint {
	result := make([]uint, num+1) // 初始化为0
	var i uint = 1
	for ; i <= num; i++ {
		result[i] = result[i&(i-1)] + 1
	}
	return result
}

func countBits3(num uint) []uint {
	result := make([]uint, num+1) // 初始化为0
	var i uint = 1
	for ; i <= num; i++ {
		result[i] = result[i>>1] + (i & 1)
	}
	return result
}
func main() {
	//fmt.Println(countBits1(10000000))
	//fmt.Println(countBits2(10000000))
	//fmt.Println(countBits3(10000000))
}
