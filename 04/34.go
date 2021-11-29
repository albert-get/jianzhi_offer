package main

import "fmt"

/**
外星语言是否排序
*/
/**
题目：有一门外星语言，它的字母表刚好包含所有的英文小写字母，只是字母表的顺序不同。给定一组单词和字母表顺序，请判断这些单词是否按照字母表的顺序排序。例如，输入一组单词["offer"，"is"，"coming"]，以及字母表顺序"zyxwvutsrqponmlkjihgfedcba"，由于字母'o'在字母表中位于'i'的前面，因此单词"offer"排在"is"的前面；同样，由于字母'i'在字母表中位于'c'的前面，因此单词"is"排在"coming"的前面。因此，这一组单词是按照字母表顺序排序的，应该输出true。
*/

func isAlienSorted(words []string, order string) bool {
	orderArray := [26]int{}
	for i, v := range order {
		orderArray[v-'a'] = i
	}
	for j := 0; j < len(words)-1; j++ {
		if !isSorted(words[j], words[j+1], orderArray) {
			return false
		}
	}
	return true
}
func isSorted(word1 string, word2 string, order [26]int) bool {
	i := 0
	for ; i < len(word1) && i < len(word2); i++ {
		ch1 := word1[i]
		ch2 := word2[i]
		if order[ch1-'a'] < order[ch2-'a'] {
			return true
		}
		if order[ch1-'a'] > order[ch2-'a'] {
			return false
		}
	}
	return i == len(word1)
}
func main() {
	fmt.Println(isAlienSorted([]string{"offer", "is", "coming"}, "zyxwvutsrqponmlkjihgfedcba"))
}
