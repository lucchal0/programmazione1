package main

import "fmt"

func èletteraValida(l rune) bool {

	ris := false

	if (l >= 'a' && l <= 'z') || (l >= 'A' && l <= 'Z') {
		ris = true
	}
	return ris
}

func èMaiuscola(l rune) bool {

	ris := false
	if l >= 'A' && l <= 'Z' {
		ris = true
	}

	return ris
}

func èVocale(l rune) bool {

	ris := false
	if l == 'a' || l == 'e' || l == 'i' || l == 'o' || l == 'u' || l == 'A' || l == 'E' || l == 'I' || l == 'O' || l == 'U' {
		ris = true
	}
	return ris
}

func main() {

	var s string
	fmt.Scan(&s)
	fmt.Println(s)

	numeroMaisc := 0
	numeroMin := 0

	for _, c := range s {
		if èletteraValida(c) {
			if èMaiuscola(c) {
				numeroMaisc++
			} else {
				numeroMin++
			}
		}
	}
	fmt.Println("maiuscole: ", numeroMaisc)
	fmt.Println("maiuscole: ", numeroMin)

	numeroVocali := 0
	numeroConsonanti := 0

	for _, c := range s {
		if èletteraValida(c) {
			if èVocale(c) {
				numeroVocali++
			} else {
				numeroConsonanti++
			}
		}
	}
	fmt.Println("vocali: ", numeroVocali)
	fmt.Println("consonanti: ", numeroConsonanti)
	

}
