package main

import "fmt"

func main() {
	var a, b float64

	fmt.Scan(&a, &b)

	var r float64

	r = -b / a

	fmt.Println("la radice è:", r)

}
