package main

import "fmt"

func main() {
	var divisore float64
	var dividendo float64

	fmt.Println("inserisci i due valori da dividere:")
	fmt.Scan(&dividendo, &divisore)

	if divisore != 0 {
		fmt.Println("quoziente =", dividendo/divisore)
	} else {
		fmt.Println("impossibile")
	}

}
