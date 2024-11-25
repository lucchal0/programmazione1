package main 

import "fmt"

func divisione_controllata(A int, b int)(ris int, ok bool){
    
    if b== 0{
        ris = 0
        ok = false
    }else{
        ris = A/B
        ok = true
 }
}

func main(){
var q, r  int
var ok  bool

    q,ok = dividisone_controllata(6,4)
    if ok {
        fmt.Println(q)
    } else {
        fmt.Println("Errore")
    }
    

}
