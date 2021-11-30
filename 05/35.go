package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

/**
最小时间差
*/
/**
题目：给定一组范围在00：00至23：59的时间，求任意两个时间之间的最小时间差。例如，输入时间数组["23：50"，"23：59"，"00：00"]，"23：59"和"00：00"之间只有1分钟的间隔，是最小的时间差。
*/

func findMinDifference(timePoints []string) int {
	if len(timePoints) > 1440 {
		return 0
	}
	minuteFlags := [1440]bool{}
	for _, point := range timePoints {
		t := strings.Split(point, "：")
		mins, _ := strconv.Atoi(t[0])
		minm, _ := strconv.Atoi(t[1])
		min := mins*60 + minm
		if minuteFlags[min] {
			return 0
		}
		minuteFlags[min] = true
	}
	return helper(minuteFlags)
}

func helper(minuteFlags [1440]bool) int {
	minDiff := len(minuteFlags) - 1
	prev := -1
	first := len(minuteFlags) - 1
	last := -1
	for i := 0; i < len(minuteFlags); i++ {
		if minuteFlags[i] {
			if prev >= 0 {
				minDiff = int(math.Min(float64(i-prev), float64(minDiff)))
			}
			prev = i
			first = int(math.Min(float64(i), float64(first)))
			last = int(math.Max(float64(i), float64(last)))
		}
	}
	minDiff = int(math.Min(float64(first+len(minuteFlags)-last), float64(minDiff)))
	return minDiff
}

func main() {
	fmt.Println(findMinDifference([]string{"23：50", "23：59", "00：00"}))
}
