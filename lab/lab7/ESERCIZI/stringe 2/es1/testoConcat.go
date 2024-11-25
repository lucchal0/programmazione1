package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Println("inserisci testo (termina con CTRL+D)")

	testoLetto := leggiTesto()
	fmt.Println(testoLetto)

}

func leggiTesto() string {

	var testo string

	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {
		testo += s.Text() + "\n"
	}

	return testo
}
