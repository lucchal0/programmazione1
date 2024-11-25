package main

import "fmt"

func main() {
	var n int
	var s int
	var R int

	fmt.Print("inserisci un numero di 3 cifre:")
	fmt.Scan(&n)

	/*
	   u = A % 10
	   A = A / 10

	   d = A % 10
	   A = A / 10

	   c= A % 10
	   A = A / 10
	*/

	for n != 0 {
		s = n % 10
		R = R + s
		n = n / 10
	}

	fmt.Println(R)

}
