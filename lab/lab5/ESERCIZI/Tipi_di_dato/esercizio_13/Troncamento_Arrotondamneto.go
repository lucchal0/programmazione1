package main

import "fmt"
import "math"

func troncamento(n float64, dopo int) float64 {
	f := math.Pow(10, float64(dopo))
	troncato := math.Trunc(n*f) / f
	return troncato
}

func arrotondamneto(n float64, dopo int) float64 {

	f := math.Pow(10, float64(dopo))
	arrotondato := math.Round(n*f) / f
	return arrotondato
}

func main() {

	var n float64
	fmt.Print("inserisci il valore:")
	fmt.Scan(&n)

	var dopoVirgola int

	fmt.Print("inserisci il numero di cifre dopo la virgola:")
	fmt.Scan(&dopoVirgola)

	fmt.Println("valore troncato = ", troncamento(n, dopoVirgola))
	fmt.Println("valore arrotondato =", arrotondamneto(n, dopoVirgola))

}
