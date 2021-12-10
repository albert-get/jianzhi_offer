package main

import (
	"fmt"
	"harry.com/jianzhi_offer/dui"
	"harry.com/jianzhi_offer/set"
	"strings"
)

/**
外星文字典
*/
/**
题目：一种外星语言的字母都是英文字母，但字母的顺序未知。给定该语言排序的单词列表，请推测可能的字母顺序。如果有多个可能的顺序，则返回任意一个。如果没有满足条件的字母顺序，则返回空字符串。例如，如果输入排序的单词列表为["ac"，"ab"，"bc"，"zc"，"zb"]，那么一个可能的字母顺序是"acbz"。
*/

func alienOrder(words []string) string {
	graph := make(map[rune]set.Set)
	inDegrees := make(map[rune]int)
	for _, word := range words {
		for _, ch := range []rune(word) {
			_, okCh := graph[ch]
			if !okCh {
				graph[ch] = *set.NewSet()
			}
			_, okChd := inDegrees[ch]
			if !okChd {
				inDegrees[ch] = 0
			}
		}
	}
	for i := 1; i < len(words); i++ {
		w1 := words[i-1]
		w2 := words[i]
		if strings.HasPrefix(w1, w2) && w1 != w2 {
			return ""
		}
		for j := 0; j < len(w1) && j < len(w2); j++ {
			ch1 := rune(w1[j])
			ch2 := rune(w2[j])
			if ch1 != ch2 {
				if s := graph[ch1]; !s.Contains(ch2) {
					s.Add(ch2)
					inDegrees[ch2] += 1
				}
				break
			}
		}
	}
	queue := dui.Queue{}
	for ch, _ := range inDegrees {
		if inDegrees[ch] == 0 {
			queue.Push(ch)
		}
	}
	var sb []rune
	for !queue.IsEmpty() {
		ch, _ := queue.Pop().(rune)
		sb = append(sb, ch)
		for next, _ := range graph[ch].M {
			next := next.(rune)
			inDegrees[next] -= 1
			if inDegrees[next] == 0 {
				queue.Push(next)
			}
		}
	}
	if len(sb) == len(graph) {
		return string(sb)
	} else {
		return ""
	}
}
func main() {
	r := alienOrder([]string{"ac", "ab", "bc", "zc", "zb"})
	fmt.Println(r)
}
