// "ceselli"
/* output: 
	"c:1"
	"e:2"
	..
	..
*/
 package main

 import (
	"fmt"
 )

 func contalettere(s string) (M map[rune]int){
	var conteggio int
	var ok bool


	M = make(map[rune]int)

	for _ , c :=range(s) {
		// se c non è mai stato visto inserisci valore 1 in M 
		// se c è già stato visto k volte inserisci k+1 in M 
		conteggio, ok = M[c]
		if (ok == false) {
			M[c] = 1
		}else{
			M[c] = conteggio +1
		}
	}

	return M 
 }

 func main(){
	var M map[rune]int 

	M = contalettere("Alessandro")

	// fmt.Println(M)

	for k, v := range M{
		fmt.Printf("%c: %d\n", k,v)
	}
 }