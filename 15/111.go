package main

import (
	"fmt"
	"harry.com/jianzhi_offer/set"
)

/**
计算除法
*/
/**
给定一个变量对数组 equations 和一个实数值数组 values 作为已知条件，其中 equations[i] = [Ai, Bi] 和 values[i] 共同表示等式 Ai / Bi = values[i] 。每个 Ai 或 Bi 是一个表示单个变量的字符串。

另有一些以数组 queries 表示的问题，其中 queries[j] = [Cj, Dj] 表示第 j 个问题，请你根据已知条件找出 Cj / Dj = ? 的结果作为答案。

返回 所有问题的答案 。如果存在某个无法确定的答案，则用 -1.0 替代这个答案。如果问题中出现了给定的已知条件中没有出现的字符串，也需要用 -1.0 替代这个答案。

*/

func calcEquation(equations [][]string, values []float32, queries [][]string) []float32 {
	graph := buildGraph(equations, values)
	results := make([]float32, len(queries))
	for i := 0; i < len(queries); i++ {
		from := queries[i][0]
		to := queries[i][1]
		_, okFrom := graph[from]
		_, okTo := graph[to]
		if !okFrom || !okTo {
			results[i] = -1
		} else {
			visited := set.NewSet()
			results[i] = dfs1(graph, from, to, visited)
		}
	}
	return results
}

func buildGraph(equations [][]string, values []float32) (graph map[string]map[string]float32) {
	graph = make(map[string]map[string]float32)
	for i := 0; i < len(equations); i++ {
		var1 := equations[i][0]
		var2 := equations[i][1]
		_, okVar1 := graph[var1]
		if okVar1 {
			graph[var1][var2] = values[i]
		} else {
			graph[var1] = make(map[string]float32)
			graph[var1][var2] = values[i]
		}
		_, okVar2 := graph[var2]
		if okVar2 {
			graph[var2][var1] = 1 / values[i]
		} else {
			graph[var2] = make(map[string]float32)
			graph[var2][var1] = 1 / values[i]
		}
	}
	return graph
}
func dfs1(graph map[string]map[string]float32, from string, to string, visited *set.Set) float32 {
	if from == to {
		return 1
	}
	visited.Add(from)
	for next, value := range graph[from] {
		if !visited.Contains(next) {
			nextValue := dfs1(graph, next, to, visited)
			if nextValue > 0 {
				return value * nextValue
			}
		}
	}
	visited.Remove(from)
	return -1
}
func main() {
	r := calcEquation([][]string{{"a", "b"}, {"b", "c"}}, []float32{2.0, 3.0}, [][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}})
	fmt.Println(r)
}
