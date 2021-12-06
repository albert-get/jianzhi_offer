package main

import "fmt"

func main() {
	a:=[]int{1,2,3,4,5}
	b:=make([]int,len(a))
	copy(b,a)
	b[0]=9
	a[0]=100
	fmt.Println(a,b)
}
