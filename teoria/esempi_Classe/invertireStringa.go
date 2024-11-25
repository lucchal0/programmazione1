package main

import "fmt"

// scrivere un sottoprogramma che inverta una stringa
// imput string s
// output string t tale cui l'ultimo elemento di s è il primo di t e via dicendo

/* func inverti(s string) (t string) {

// 	t = ""

// 	for i := 0; i < len(s); i++ {
// 		t[i] = s[len(s)-1-i]
// 	}

// 	return
 }
*/

func invertiS(s string) (t string) {
	for _, c := range s {
		t = string(c) + t
	}
	return
}

func main() {

	s := "ciao"
	fmt.Println(invertiS(s))
}
