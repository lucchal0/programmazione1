package main

import "fmt"

func test(x int) (y int, z int){

	z= x + 2
	y= x + 1
	return
}

func main() {
	var a, b int
	a, b = test(10)
	fmt.Println(a, b)
}
