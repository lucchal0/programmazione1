// tabelline

package main

import "fmt"

func tabellina(n int) {
	var i int
	for i = 0; i <= 10; i++ {
		fmt.Print(" ", n*i)
	}
	return
}

func main() {
	var Numero int

	fmt.Print("inserisci un numero: ")
	fmt.Scan(&Numero)

	tabellina(Numero)
}
