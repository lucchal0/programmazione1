package main

import "fmt"

// CONTO ALLA ROVESCIA

func main() {

	var i int
	var n int

	fmt.Scan(&n)

	for i = n; i >= 0; i = i - 1 {
		fmt.Println("contatore:", i)
	}

}
