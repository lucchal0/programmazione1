package main

import "fmt"
import "fmt"
import "os"

/* capire se le stringe è il prefisso di un'altra
------------------------------------------------------------------
   input: stringa e prefisso di tipo stringa
   output : true  -> prefisso (bool)
		    false -> non corrisponde (bool)
	dove per prefisso intendo che tutti i caratteri di s compaiono nello stesso ordine al inizio di t

	str: "animale", prefix:"animalesco" -> true

	str: "piantare", t: piata -> false
	str: "piantare", t: "" -> false

------------------------------------------------------------------
*/

func prefisso(str, prefix string) bool {


	if len(prefix) > len(str) {
		return false
	} else {
/*
		// Confronta carattere per carattere
		for i := 0; i < len(prefix); i++ {
			if str[i] != prefix[i] {
				return false
			}
		}
	}
	return true

	---------------------------------------------------------------------------

	if s == t[:len(s)]{
		return true
	} else{
		return false 
	}
	*/

	return str== prefix[:len(str)]
	}
}

func main() {

	// stringa 1 e stringa 2
	var str1, pref string

	fmt.Print("inserisci prima stringa: ")
	fmt.Scan(&str1)

	fmt.Print("inserisci prefisso: ")
	fmt.Scan(&pref)

	if prefisso(str1, pref) {
		// len(str1) > len(str2)
		fmt.Printf("'%s' è un prefisso di '%s'\n", pref, str1)
	} else {
		// str1 != str2
		fmt.Println("Nessuna delle due stringhe è prefisso dell'altra.")
	}

}
