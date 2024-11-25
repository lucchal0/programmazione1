package main

import "fmt"

func main() {

	var n float64

	fmt.Print("inserisci il numero da arrontondare:")
	fmt.Scan(&n)

	arrotondamento := (float64(int((n * 100) + 0.5))) / 100

	fmt.Print("Valore arrontondato:", arrotondamento)

}
