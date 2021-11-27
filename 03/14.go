package main

import "fmt"

/**
字符串中的变位词
*/
/**
题目：输入字符串s1和s2，如何判断字符串s2中是否包含字符串s1的某个变位词？如果字符串s2中包含字符串s1的某个变位词，则字符串s1至少有一个变位词是字符串s2的子字符串。假设两个字符串中只包含英文小写字母。例如，字符串s1为"ac"，字符串s2为"dgcaf"，由于字符串s2中包含字符串s1的变位词"ca"，因此输出为true。如果字符串s1为"ab"，字符串s2为"dgcaf"，则输出为false。
*/

func checkInclusion(c string, sp string) bool {
	if len(sp) < len(c) {
		return false
	}
	counts := [26]int{}
	for i, v := range c {
		counts[v-'a']++
		counts[sp[i]-'a']--
	}
	if areAllZero(counts) {
		return true
	}
	for i := len(c); i < len(sp); i++ {
		counts[sp[i]-'a']--
		counts[sp[i-len(c)]-'a']++
		if areAllZero(counts) {
			return true
		}
	}
	return false

}
func areAllZero(counts [26]int) bool {
	for _, v := range counts {
		if v != 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(checkInclusion("ac", "dgcaf"))
}
