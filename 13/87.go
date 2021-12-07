package main

import (
	"fmt"
	"strconv"
)

/**
恢复IP地址
*/
/**
题目：输入一个只包含数字的字符串，请列出所有可能恢复出来的IP地址。例如，输入字符串"10203040"，可能恢复出3个IP地址，分别为"10.20.30.40"、"102.0.30.40"和"10.203.0.40"。
*/

func restoreIpAddresses(s string) []string {
	result := new([]string)
	helper8(s, 0, 0, "", "", result)
	return *result
}
func helper8(s string, i int, segI int, seg string, ip string, result *[]string) {
	if i == len(s) && segI == 3 && isValidSeg(seg) {
		*result = append(*result, ip+seg)
	} else if i < len(s) && segI <= 3 {
		ch := string(s[i])
		if isValidSeg(seg + ch) {
			helper8(s, i+1, segI, seg+ch, ip, result)
		}
		if len(seg) > 0 && segI < 3 {
			helper8(s, i+1, segI+1, ""+ch, ip+seg+".", result)
		}

	}
}
func isValidSeg(seg string) bool {
	v, _ := strconv.Atoi(seg)
	return v <= 255 && (seg == "0" || string(seg[0]) != "0")
}
func main() {
	r := restoreIpAddresses("10203040")
	fmt.Println(r)
}
