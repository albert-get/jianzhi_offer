package main

import "fmt"

/**
所有路径
*/
/**
题目：一个有向无环图由n
个节点（标号从0到n
-1，n
≥2）组成，请找出从节点0到节点n
-1的所有路径。图用一个数组graph表示，数组的graph[i
]包含所有从节点i
能直接到达的节点。例如，输入数组graph为[[1，2]，[3]，[3]，[]]，则输出两条从节点0到节点3的路径，分别为0→1→3和0→2→3，
*/

func allPathsSourceTarget(graph [][]int) [][]int {
	result := new([][]int)
	path := new([]int)
	dfs(0, graph, path, result)
	return *result
}
func dfs(source int, graph [][]int, path *[]int, result *[][]int) {
	*path = append(*path, source)
	if source == len(graph)-1 {
		*result = append(*result, *path)
	} else {
		for _, next := range graph[source] {
			dfs(next, graph, path, result)
		}
	}
	t := *path
	*path = append([]int{}, t[:len(*path)-1]...)
}
func main() {
	p := allPathsSourceTarget([][]int{{1, 2}, {3}, {3}, {}})
	fmt.Println(p)
}
