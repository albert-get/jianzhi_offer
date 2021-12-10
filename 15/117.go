package main

import "fmt"

/**
相似的字符串
*/
/**
题目：如果交换字符串X中的两个字符就能得到字符串Y，那么两个字符串X和Y相似。例如，字符串"tars"和"rats"相似（交换下标为0和2的两个字符）、字符串"rats"和"arts"相似（交换下标为0和1的字符），但字符串"star"和"tars"不相似。
*/

func numSimilarGroups(A []string) int {
	fathers := make([]int, len(A))
	for i := 0; i < len(fathers); i++ {
		fathers[i] = i
	}
	groups := len(A)
	for i := 0; i < len(A); i++ {
		for j := i + 1; j < len(A); j++ {
			if similar(A[i], A[j]) && union1(&fathers, i, j) {
				groups--
			}
		}
	}
	return groups
}
func findFather1(fathers *[]int, i int) int {
	t := *fathers
	if t[i] != i {
		findFather1(fathers, t[i])
	}
	return t[i]
}
func union1(fathers *[]int, i int, j int) bool {
	t := *fathers
	fatherOfI := findFather1(fathers, i)
	fatherOfJ := findFather1(fathers, j)
	if fatherOfI != fatherOfJ {
		t[fatherOfI] = fatherOfJ
		return true
	}
	return false
}
func similar(str1 string, str2 string) bool {
	diffcount := 0
	for i := 0; i < len(str1); i++ {
		if str1[i] != str2[i] {
			diffcount++
		}
	}
	return diffcount <= 2
}
func main() {
	f := numSimilarGroups([]string{"tars", "rats", "arts", "star"})
	fmt.Println(f)
}
