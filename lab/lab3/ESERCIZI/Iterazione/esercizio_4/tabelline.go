package main

import "fmt"

func main(){

var n int 

    fmt.Print("inserisci un numero: ")
    fmt.Scan(&n)

var i int 

    for i=1; i<=10; i++{

        fmt.Println(i," X ", n,"=", i * n) // stampo e moltiplico 
        
    }

}
