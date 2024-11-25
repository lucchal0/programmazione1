/*
Scrivere un programma che legga da standard input un numero intero n e, come mostrato nell'Esempio di
esecuzione, stampi a video un triangolo rettangolo con base e altezza di lunghezza n utilizzando il carattere
* (asterisco).
Se n è negativo o nullo, anziché stampare il triangolo il programma deve stampare un messaggio d'errore.
Nota: Si utilizzi un solo ciclo for ed una variabile di tipo string .

$ go run triangolo.go
5
*
**
***
****
*****
*/
package main

import "fmt"

func main(){
	var n int
	fmt.Scan(&n)

	if (n<0){
		fmt.Println("Dimensione non sufficiente")
	}else{
		s:= " "
		for i:=0; i<n; i++{
			s = s + "*"
			fmt.Println(s)
	}
	}
}