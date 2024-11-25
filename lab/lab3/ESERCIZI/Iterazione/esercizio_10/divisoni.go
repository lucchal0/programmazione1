package main

import "fmt"

func main(){
    var N int 
    var i int 

    fmt.Print("inserisci un numero:")
    fmt.Scan(&N)
    
    for i=1; i<N; i++{
        if N%i==0{
            fmt.Print(i, " ")
        }
    }


}
