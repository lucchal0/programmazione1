package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type frazione struct {
	numeratore, denominatore int
}

func LeggiFrazioni() []frazione {
	var frazioni []frazione
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Inserisci numeratore e denominatore delle frazioni (CTRL+D per terminare):")
	for scanner.Scan() {
		linea := scanner.Text()
		nums := strings.Fields(linea)

		n, _ := strconv.Atoi(nums[0])
		d, _ := strconv.Atoi(nums[1])

		frazioni = append(frazioni, frazione{numeratore: n, denominatore: d})
	}

	return frazioni
}

// Funzione per moltiplicare due frazioni
func Moltiplica(f1, f2 frazione) *frazione {
	return &frazione{
		numeratore:   f1.numeratore * f2.numeratore,
		denominatore: f1.denominatore * f2.denominatore,
	}
}

// Funzione per moltiplicare una slice di frazioni
func MoltiplicaN(fN []frazione) *frazione {
	if len(fN) == 0 {
		return &frazione{numeratore: 1, denominatore: 1} // Frazione neutra
	}

	prodotto := fN[0] // Inizializza il prodotto con la prima frazione
	for i := 1; i < len(fN); i++ {
		prodotto = *Moltiplica(prodotto, fN[i])
	}
	return &prodotto
}

// ------------------------------
func mcd(a, b int) int {
	if b == 0 {
		return a
	}
	return mcd(b, a%b)
}

// Funzione che riduce ai minimi termini una frazione.
func (f *frazione) Riduci() {
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

//------------------------------

func main() {
	// Leggere le frazioni da standard input
	frazioni := LeggiFrazioni()

	prodotto := MoltiplicaN(frazioni)

	prodotto.Riduci()

	// Stampare il risultato
	fmt.Printf("Prodotto: %d/%d\n", prodotto.numeratore, prodotto.denominatore)

}
