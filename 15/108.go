package main

import (
	"fmt"
	"harry.com/jianzhi_offer/dui"
	"harry.com/jianzhi_offer/set"
)

/**
单词演变
*/
/**
题目：输入两个长度相同但内容不同的单词（beginWord和endWord）和一个单词列表，求从beginWord到endWord的演变序列的最短长度，要求每步只能改变单词中的一个字母，并且演变过程中每步得到的单词都必须在给定的单词列表中。如果不能从beginWord演变到endWord，则返回0。假设所有单词只包含英文小写字母。
*/

func ladderLength(beginWord string, endWord string, wordList []string) int {
	queue1 := dui.Queue{}
	queue2 := dui.Queue{}
	notVisited := set.NewSet()
	for _, w := range wordList {
		notVisited.Add(w)
	}
	length := 1
	queue1.Push(beginWord)
	for !queue1.IsEmpty() {
		word, _ := queue1.Pop().(string)
		if word == endWord {
			return length
		}
		neighbors := getNeighbors(word)
		for _, neighbor := range neighbors {
			if notVisited.Contains(neighbor) {
				queue2.Push(neighbor)
				notVisited.Remove(neighbor)
			}
		}
		if queue1.IsEmpty() {
			length++
			queue1 = queue2
			queue2 = dui.Queue{}
		}
	}
	return 0
}
func getNeighbors(word string) []string {
	var neighbors []string
	chars := []rune(word)
	for i := 0; i < len(chars); i++ {
		old := chars[i]
		for ch := 'a'; ch <= 'z'; ch++ {
			if old != ch {
				chars[i] = ch
				neighbors = append(neighbors, string(chars))
			}
		}
		chars[i] = old
	}
	return neighbors
}

func ladderLength1(beginWord string, endWord string, wordList []string) int {
	notVisited := set.NewSet()
	for _, w := range wordList {
		notVisited.Add(w)
	}
	if !notVisited.Contains(endWord) {
		return 0
	}
	set1 := set.NewSet()
	set2 := set.NewSet()
	length := 2
	set1.Add(beginWord)
	set2.Add(endWord)
	for set1.Size() > 0 && set2.Size() > 0 {
		if set1.Size() < set2.Size() {
			temp := set1
			set1 = set2
			set2 = temp
		}
		set3 := set.NewSet()
		for k, _ := range set1.M {
			word, _ := k.(string)
			neighbors := getNeighbors(word)
			for _, neighbor := range neighbors {
				if set2.Contains(neighbor) {
					return length
				}
				if notVisited.Contains(neighbor) {
					set3.Add(neighbor)
					notVisited.Remove(neighbor)
				}
			}
		}
		length++
		set1 = set3
	}
	return 0
}
func main() {
	//r := ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"})
	r := ladderLength1("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"})
	fmt.Println(r)
}
