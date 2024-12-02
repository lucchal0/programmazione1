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
		if unicode.IsLetter(lettera){
		m[lettera] += 1
		}
	}

	return m

}

func main() {

	//fmt.Println(leggiTesto())

	m := occorenze(leggiTesto())

	for k, v := range m {
		fmt.Printf("%c -> %d\n", k, v)
	}

}
