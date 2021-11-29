package main

import (
	"fmt"
	"sort"
)

/**
变位词组
*/
/**
题目：给定一组单词，请将它们按照变位词分组。例如，输入一组单词["eat"，"tea"，"tan"，"ate"，"nat"，"bat"]，这组单词可以分成3组，分别是["eat"，"tea"，"ate"]、["tan"，"nat"]和["bat"]。假设单词中只包含英文小写字母。
*/

func groupAnagrams(strs []string) (res [][]string) {
	hash := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101}
	groups := make(map[int64][]string)
	for _, str := range strs {
		var wordHash int64 = 1
		for _, s := range str {
			wordHash *= int64(hash[s-'a'])
		}
		groups[wordHash] = append(groups[wordHash], str)
	}
	for _, v := range groups {
		res = append(res, v)
	}
	return res
}

type arrS []rune

func (m arrS) Len() int {
	return len(m)
}
func (m arrS) Less(i, j int) bool {
	return m[i] < m[j]
}
func (m arrS) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func groupAnagrams2(strs []string) (res [][]string) {
	groups := make(map[string][]string)
	for _, str := range strs {
		charArray := arrS(str)
		sort.Sort(charArray)
		sorted := string(charArray)
		groups[sorted] = append(groups[sorted], str)
	}
	for _, v := range groups {
		res = append(res, v)
	}
	return res
}
func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	fmt.Println(groupAnagrams2([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}
