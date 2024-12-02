package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	// slice di interi
	numeri := make([]int, n)

	fmt.Println("inserisci", n, "numeri:")
	for i := 0; i < n; i++ {
		fmt.Scan(&numeri[i])
	}

	// invertire numeri
	fmt.Println("Numeri in ordine inverso:")
	for i := n - 1; i >= 0; i-- {
		fmt.Print(numeri[i], " ")
	}
}
