package main

import "fmt"

func main() {

	var A, R int

	fmt.Print("valore da calcolare la radice:")
	fmt.Scan(&A)

	for R*R < A {
		R++
	}

	fmt.Println("la radice di", A, "è", R)

}
