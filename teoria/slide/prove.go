package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var lines []string // Slice per memorizzare le righe di testo

	fmt.Println("Inserisci un testo (premi Ctrl+D per terminare):")

	// Leggi tutte le righe finché non viene premuto Ctrl+D (EOF)
	for scanner.Scan() {
		lines = append(lines, scanner.Text()) // Aggiungi ogni riga allo slice
	}

	// Verifica se c'è stato un errore durante la lettura
	if err := scanner.Err(); err != nil {
		fmt.Println("Errore durante la lettura:", err)
	}

	// Stampa le righe in ordine inverso
	fmt.Println("\nTesto al contrario:")
	for i := len(lines) - 1; i >= 0; i-- {
		fmt.Println(lines[i])
	}
}
