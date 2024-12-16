package main

import "fmt"

// Definizione dell'interfaccia  Forma
// Una  Forma  deve implementare due metodi:  Area  e  Perimetro
type Forma interface {
	Area() float64
	Perimetro() float64
}

// Definizione della struttura  rettangolo
type rettangolo struct {
	larghezza, altezza float64
}

// Implementazione del metodo  Area  per il tipo  rettangolo
func (r rettangolo) Area() float64 {
	return r.altezza * r.larghezza
}

// Implementazione del metodo  Perimetro  per il tipo  rettangolo
func (r rettangolo) Perimetro() float64 {
	return 2 * (r.larghezza + r.altezza)
}

func main() {
	r := rettangolo{larghezza: 10.0, altezza: 20.0}

	// Dichiarazione di una variabile  f  di tipo  Forma  e assegnazione di  r .
	// Questo è possibile perché  rettangolo  implementa l'interfaccia  Forma .
	var f Forma = r

	// Stampa della rappresentazione della variabile  f
	fmt.Println("f = r ->", f)

	// Chiamata dei metodi  Area  e  Perimetro  attraverso l'interfaccia  Forma .
	// Questi metodi sono definiti su  rettangolo  e vengono eseguiti su  r .
	fmt.Println("area: ", f.Area())
	fmt.Println("perimetro: ", f.Perimetro())
}
