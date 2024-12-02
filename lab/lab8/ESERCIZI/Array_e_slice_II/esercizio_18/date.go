package main

import (
	"bufio"
	"fmt"
	"os"
)

func leggiTesto() (date []string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		riga := scanner.Text()
		if riga != "" {
			date = append(date, riga)
		}
	}
	return
}

func formatta(data string) string {
	dataRune := []rune(data)

	switch len(dataRune) {
	case 8:
		// Esempio: "YYYYMMDD" -> "YYYY-MM-0D"
		return string(append(dataRune[:4], '-', dataRune[4], '-', '0', dataRune[7]))
	case 9:
		if dataRune[6] == '/' {
			// Esempio: "YYYY/M/DD" -> "YYYY-M-DD"
			return string(append(dataRune[:4], '-', dataRune[5], '-', dataRune[7], dataRune[8]))
		} else {
			// Esempio: "YYYYMM/D" -> "YYYY-MM-0D"
			return string(append(dataRune[:4], '-', dataRune[4], dataRune[5], '-', '0', dataRune[8]))
		}
	default:
		if len(dataRune) >= 10 {
			// Se la stringa è già ben formattata
			dataRune[4] = '-'
			dataRune[7] = '-'
		}
	}
	return string(dataRune)
}

func main() {
	date := leggiTesto()

	s := "Ciao ⌘"
runes := []rune(s)
fmt.Println(runes)// Output: valori Unicode di ogni carattere

	for _, data := range date {
		fmt.Println(formatta(data))
	}
}
