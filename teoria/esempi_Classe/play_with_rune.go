package main

import "fmt"

/* controllo se una rune in input è un carattere minhusolo del alfabbeto latino oppure no 
input R rune
Output: true se R .... false altrimenti
*/

func isLowerCase(R rune) bool{

	// tra 97 e 122 inclusi
	if R >= 'a' && R <= 'z'{
		retrun true
	}else{
		return false
	}
}

func main() {
	var carattere rune

	// 'A' costante di tipo rune
	carattere = 'A'
	fmt.Println(carattere)


}
