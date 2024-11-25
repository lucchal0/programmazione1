package main

import "fmt"

func main() {

	var n float64

	fmt.Print("inserisci valore da troncare:")
	fmt.Scan(&n)

	troncato := (float64(int(n * 100))) / 100
	fmt.Print("valore torncato:", troncato)

}
