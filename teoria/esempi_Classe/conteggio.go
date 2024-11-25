package main

import "fmt"

func main() {
	//scrivere un programma preso in input un numero non negativo stampa tutti i numeri interi tra 0 e n in input compreso

	var n int
	var i int

	fmt.Print("inserisci un numero: ")
	fmt.Scan(&n)

	/*i = 0
	  for i<= n {
	      fmt("contatore: ",i)
	      i++
	  }*/

	for i = 0; i < n; i++ {
		fmt.Println("output:", i)
	}

}
