package main

func testmap() (graph map[string]map[string]float32) {
	graph["ok"]["ok"]=1
	return graph
}
func main() {
	testmap()
}
