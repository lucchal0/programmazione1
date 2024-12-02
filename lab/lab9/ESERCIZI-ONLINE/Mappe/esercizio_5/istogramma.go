package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func leggiTesto() string {
	var riga string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		riga += scanner.Text() + "\n"
	}
	return riga[:len(riga)]
}

func occorenze(s string) map[rune]int {

	m := map[rune]int{}
	for _, lettera := range s {
		if unicode.IsLetter(lettera) {
			m[lettera] += 1
		}
	}

	return m

}

func stampaoccorenze(m map[rune]int) {
	fmt.Println("istogramma:")
	for k, v := range m {
		fmt.Printf("%c: %s\n", k, creaBarra(v))
	}
}

func creaBarra(k int) string {
	s := ""

	for i := 1; i <= k; i++ {
		s += "*"
	}
	return s
}

func main() {

	//fmt.Println(leggiTesto())

	stampaoccorenze(occorenze(leggiTesto()))
}
