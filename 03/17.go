package main

import (
	"fmt"
	"math"
)

/**
包含所有字符的最短字符串
*/
/**
题目：输入两个字符串s和t，请找出字符串s中包含字符串t的所有字符的最短子字符串。例如，输入的字符串s为"ADDBANCAD"，字符串t为"ABC"，则字符串s中包含字符'A'、'B'和'C'的最短子字符串是"BANC"。如果不存在符合条件的子字符串，则返回空字符串""。如果存在多个符合条件的子字符串，则返回任意一个。
*/

func minWindow(s string, t string) string {
	charToCount := make(map[string]int)
	for _, v := range t {
		charToCount[string(v)]++
	}
	count := len(charToCount)
	start := 0
	end := 0
	minStart := 0
	minEnd := 0
	minLength := math.MaxInt
	for end < len(s) || (count == 0 && end == len(s)) {
		if count > 0 {
			/**
			end每右移，如果有就要抵消一次
			还原的时候如果从有变无，说明长度要减一
			*/
			endCh := s[end]
			_, ok := charToCount[string(endCh)]
			if ok {
				charToCount[string(endCh)]--
				if charToCount[string(endCh)] == 0 {
					count--
				}
			}
			end++
		} else {
			if end-start < minLength {
				minLength = end - start
				minStart = start
				minEnd = end
			}
			/**
			start每右移，如果有就要还原一次
			还原的时候如果从无变有，说明长度要加一
			*/
			startCh := s[start]
			_, ok := charToCount[string(startCh)]
			if ok {
				charToCount[string(startCh)]++
				if charToCount[string(startCh)] == 1 {
					count++
				}
			}
			start++
		}
	}
	if minLength < math.MaxInt {
		return s[minStart:minEnd]
	} else {
		return ""
	}
}
func main() {
	fmt.Println(minWindow("ADDBANCAD", "ABC"))
}
