package main 

import "fmt"
import "math/rand"
import "time"

func main(){
    // inzializza 
    rand.Seed(int64(time.Now().Nanosecond()))
    // genere tra 0 e 99
    var random int = rand.Intn(100)
    var count int = 1
    var n int

    //fmt.Print(random)

    for{
        fmt.Print("tentativo n°", count,":")
        fmt.Scan(&n)

        if n > random {
            fmt.Println("Troppo alto riprova")
            count ++
        }else{
            if n < random {
            fmt.Println("Troppo basso riprova")
            count ++
            } 
        }
        if n == random{
            fmt.Println("Trovato!!")
            break
        }
    }

    fmt.Println("hai indovinato in ", count, "tentativi!")
}
