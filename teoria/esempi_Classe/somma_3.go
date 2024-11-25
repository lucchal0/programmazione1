//somma un numero arbitrario di numeri e fermarsi quando valore è 0

package main

import "fmt"

func main(){
var N int
var A int
var somma int 

    fmt.Scan(&N)

    somma = 0
    
    for {
        fmt.Scan(&A)
        if A == 0{
            break
        }
        somma = somma + A
    }

}
