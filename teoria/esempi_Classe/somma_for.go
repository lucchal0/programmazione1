//leggere numeri interi e stampare la somma

package main

import "fmt"

func main() {
	var n int
	var somma int
	var N int
	var i int

	fmt.Scan(&N)

	for i = 0; i < N; i++ {
		fmt.Print("inserisci il numero : ")
		fmt.Scan(&n)
		somma = somma + n
	}

	fmt.Print("somma =", somma)

}
