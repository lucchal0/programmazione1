package main 

import "fmt"

func main(){

    var n float64
    var media float64
    var count float64

    fmt.Print("inserisci una sequenza di numeri:" )
    for{
        fmt.Scan(&n)
        if n <= 0.0{
            break
        }else{
            media = media + n
            count ++
        }
    }

    fmt.Print("media: ",media/count)






}
