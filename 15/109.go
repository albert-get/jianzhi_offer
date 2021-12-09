package main

import (
	"fmt"
	"harry.com/jianzhi_offer/dui"
	"harry.com/jianzhi_offer/set"
)

/**
开密码锁
*/
/**
题目：一个密码锁由4个环形转轮组成，每个转轮由0～9这10个数字组成。每次可以上下拨动一个转轮，如可以将一个转轮从0拨到1，也可以从0拨到9。密码锁有若干死锁状态，一旦4个转轮被拨到某个死锁状态，这个锁就不可能打开。密码锁的状态可以用一个长度为4的字符串表示，字符串中的每个字符对应某个转轮上的数字。输入密码锁的密码和它的所有死锁状态，请问至少需要拨动转轮多少次才能从起始状态"0000"开始打开这个密码锁？如果锁不可能打开，则返回-1。
*/

func openLock(deadens []string, target string) int {
	dead := set.NewSet()
	for _, d := range deadens {
		dead.Add(d)
	}
	visited := set.NewSet()
	init := "0000"
	if dead.Contains(init) || dead.Contains(target) {
		return -1
	}
	queue1 := dui.Queue{}
	queue2 := dui.Queue{}
	steps := 0
	queue1.Push(init)
	visited.Add(init)
	for !queue1.IsEmpty() {
		cur, _ := queue1.Pop().(string)
		if cur == target {
			return steps
		}
		nexts := getNeighbors1(cur)
		for _, next := range nexts {
			if !dead.Contains(next) && !visited.Contains(next) {
				queue2.Push(next)
				visited.Add(next)
			}
		}
		if queue1.IsEmpty() {
			steps++
			queue1 = queue2
			queue2 = dui.Queue{}
		}
	}
	return -1
}
func getNeighbors1(cur string) []string {
	var nexts []string
	for i := 0; i < len(cur); i++ {
		ch := cur[i]
		var newCh rune
		if ch == '0' {
			newCh = '9'
		} else {
			newCh = rune(ch - 1)
		}
		builder := []rune(cur)
		builder[i] = newCh
		nexts = append(nexts, string(builder))
		if ch == '9' {
			newCh = '0'
		} else {
			newCh = rune(ch + 1)
		}
		builder[i] = newCh
		nexts = append(nexts, string(builder))
	}
	return nexts
}
func main() {
	r := openLock([]string{"0102", "0201"}, "0202")
	fmt.Println(r)
}
