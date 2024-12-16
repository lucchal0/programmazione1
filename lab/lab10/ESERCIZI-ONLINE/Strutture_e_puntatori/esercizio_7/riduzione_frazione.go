package main

import (
	"fmt"
	"os"
	"strconv"
)

type frazione struct {
	numeratore, denominatore int
}

// Funzione che calcola il massimo comune divisore (MCD) usando l'algoritmo di Euclide.
func mcd(a, b int) int {
	if b == 0 {
		return a
	}
	return mcd(b, a%b)
}

// Funzione che riduce ai minimi termini una frazione.
func Riduci(f *frazione) {
	if f.denominatore == 0 {
		fmt.Println("Errore: il denominatore non può essere zero!")
		return
	}

	divisore := mcd(f.numeratore, f.denominatore)
	f.numeratore /= divisore
	f.denominatore /= divisore

	// Assicuriamoci che il denominatore sia positivo.
	if f.denominatore < 0 {
		f.numeratore = -f.numeratore
		f.denominatore = -f.denominatore
	}
}

func main() {

	// Conversione degli argomenti in interi.
	n, _ := strconv.Atoi(os.Args[1])
	d, _ := strconv.Atoi(os.Args[2])

	// Creazione della frazione.
	f := frazione{n, d}

	Riduci(&f)

	fmt.Printf("%d/%d\n", f.numeratore, f.denominatore)
}
