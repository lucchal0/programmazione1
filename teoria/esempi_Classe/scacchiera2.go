package main

import (
	"fmt"
	"os"
	"strconv"
)

func inputDimensione(argomenti []string) (N, M int, tutto_ok bool) {
	if len(argomenti) < 3 {
		return 0, 0, false // Non ci sono abbastanza argomenti
	}

	var err error
	N, err = strconv.Atoi(argomenti[1])
	if err != nil {
		return 0, 0, false // Errore nella conversione di N
	}
	M, err = strconv.Atoi(argomenti[2])
	if err != nil {
		return 0, 0, false // Errore nella conversione di M
	}

	if N < 0 || M < 0 {
		return 0, 0, false // Dimensioni negative
	}
	return N, M, true
}

// input: numero di righe N e il numero di colonne M
// output : S [][]bool, scacchiera di NxM, false = nero, true = bianco
func creaScacchiera(N, M int) (S [][]bool) {
	S = make([][]bool, N)
	for i := 0; i < N; i++ {
		S[i] = make([]bool, M)
		for j := 0; j < M; j++ { // Corretto da i++ a j++
			if (i % 2) == (j % 2) {
				S[i][j] = true // Bianco
			} else {
				S[i][j] = false // Nero
			}
		}
	}
	return
}

func stampascacchiera(S [][]bool, blackChar, whiteChar string) {
	for _, row := range S {
		for _, cell := range row {
			if cell {
				fmt.Print(string(whiteChar))
			} else {
				fmt.Print(string(blackChar))
			}
		}
		fmt.Println()
	}
}

func main() {
	var N, M int
	var ok bool
	var T [][]bool

	// Input
	N, M, ok = inputDimensione(os.Args)
	if !ok {
		fmt.Println("Errore: dimensioni non valide")
		return
	}

	// Creare scacchiera
	T = creaScacchiera(N, M)

	// Stampare
	stampascacchiera(T, " * ", " - ")
}
