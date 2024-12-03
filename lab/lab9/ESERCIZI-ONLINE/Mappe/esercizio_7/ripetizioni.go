package main

import (
	"fmt"
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

func separa(s string) []string{
	sub:=separaSub(s)
	parole:= []string{}

	for _, sub:= range sub{
		if èparola(sub){
			parole = append(parole, sub)
		}
	}
	return parole
}

func èparola(sub string) bool{
	
	for _, lettera := range sub{
		if !unicode.IsLetter(lettera){
			return false
		}
	}
	return true
}

func separaSub(s string) []string{
	sott:=[]string{}

	 predenteNonècarattere := true
	ss:=""

	for _, c:= range s{
		if !unicode.IsSpace(c) &&  predenteNonècarattere{
			sott= append(sott, ss)
		} 
		ss= string(c)
		predenteNonècarattere= false
	}

	return sott
}

func main(){

	s:=leggiTesto()
	parole:= separa(s)

}