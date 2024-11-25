package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {

	fmt.Println("Inserisci testo (termina con CTRL+D):")

	// Leggi il testo dall'input standard
	testoLetto := leggiTesto()

	// Stampa il testo letto
	fmt.Println("Testo inserito:\n", testoLetto)

	// Calcola statistiche su parole e caratteri
	numeroParole, numeroCar := èUnaparola(testoLetto)

	media := float32(numeroCar) / float32(numeroParole)

	// Stampa i risultati
	fmt.Printf("Numero di parole: %d\n", numeroParole)
	fmt.Printf("media: %f",media)
}

func leggiTesto() string {
	var testo string
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		testo += scanner.Text() + "\n"
	}

	return testo
}

// Funzione per calcolare il numero di parole e caratteri in una stringa
func èUnaparola(s string) (int, int) {
	var numeroParole, numeroCaratteri int
	var ultimoLettoÈCarattere bool = false

	for _, c := range s {
		if unicode.IsLetter(c) {
			if !ultimoLettoÈCarattere {
				numeroParole++
			}
			ultimoLettoÈCarattere = true
			numeroCaratteri++
		} else {
			ultimoLettoÈCarattere = false
		}
	}
	return numeroParole, numeroCaratteri
}
