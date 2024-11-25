package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
    fmt.Println("testo invertito: \n",testo())
}

func testo() string{

	var newStr string
    fmt.Println("inserisci testo formato da più righe (termine con riga vuota)")
    s := bufio.NewScanner(os.Stdin)
	for {
		s.Scan()
		if s.Text() == "" {
			break
		}
		newStr = newStr + s.Text() +"\n"
	}
    return  leggi_Contrario(newStr)
}

func leggi_Contrario(s string)(string){

    var str string
    for i:=len(s)-1; i>=0; i--{
        str += string(s[i])
    } 
    return str

}
