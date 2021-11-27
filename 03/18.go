package main

import (
	"fmt"
	"unicode"
)

/**
有效的回文
*/
/**
题目：给定一个字符串，请判断它是不是回文。假设只需要考虑字母和数字字符，并忽略大小写。例如，"Was it a cat I saw？"是一个回文字符串，而"race a car"不是回文字符串。
*/
func isPalindrome(s string) bool {
	i := 0
	j := len(s) - 1
	for i < j {
		ch1 := s[i]
		ch2 := s[j]
		if !unicode.IsLetter(rune(ch1)) || !unicode.IsDigit(rune(ch1)) {
			i++
		} else if !unicode.IsLetter(rune(ch2)) || !unicode.IsDigit(rune(ch2)) {
			j--
		} else {
			ch1 = uint8(unicode.ToLower(rune(ch1)))
			ch2 = uint8(unicode.ToUpper(rune(ch2)))
			if ch1 != ch2 {
				return false
			}
			i++
			j--
		}
	}
	return true
}
func main() {
	fmt.Println(isPalindrome("Was it a cat I saw？"))
}
