package main 

import "fmt"
import "math"

func main(){
    
    var a,b,c float64
    var delta, x1 ,x2 float64

    fmt.Print("inserisci i dati dell'equazione a b c:" )
    fmt.Scan(&a, &b, &c)

    delta = (b*b) - 4(a*c)

    if delta > 0{
        x1 = (- b - math.Sqrt(delta))/2*a
        x2 = (- b + math.Sqrt(delta))/2*a
        fmt.Println("x1 =",x1)
        fmt.Prntln("x2 =",x2)
    } 
    else { 
        if delta == 0{
          x1= -b/(2*a)
          fmt.Print("x1 è uguale a x2=",x1)
          }else{
        fmt.Println("equazione impossibile siccome il delta è negativo")
    }}

}
