package main 

import "fmt"

func main(){

    var n int 
        
        fmt.Println("inserisci un numero: ")
        fmt.Scan(&n)

        fmt.Println("inserisci",n,"numeri")

    var i int = 1
    var numero int

    var somma int
    var positivo int
    var negativo int 
    var zero int


    var minimo int
    var massimo int

        for i<=n, i++{
            // scann
            fmt.Scan(&numero)
            //somma
            somma = somma + numero
            // minimo
            if i == 1 || minimo > massimo{
                    minimo = numero
            }

            //n>0
            if numero > 0 {
                positivo ++
            }else {
                if numero < 0{
                    negativo ++ 
                } else {
                    zero ++
                }}
        }

// stampa finale 

        fmt.Println("Somma =", somma)
        fmt.Println("Valore minimo =", minimo)
        fmt.Println("Interi > 0 =", positivo)
        fmt.Println("Interi < 0 =", negativo)
        fmt.Println("Interi = 0 =", zero) 

}
