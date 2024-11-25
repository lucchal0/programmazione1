/*
Calcolare il perimetro di un triangolo 

Specifiche: 
- l'utente inserisce in input le cordinate dei vertici in un piano cartesiano
- il programma stampa un singolo valore, che è il perimentro del triangolo ottenuto colegando i vertici

Processo risolutivo: 
 - dai vertici clacolo la lunghezza dei lati 
 - dalle lunghezze dei lati calcolo il perimetro 

 Sottoproblema: 
 calcolare la lunghezza di un lato 
 - input : vertici A, B
 - output: distanza tra A e B 
*/
package main

import "fmt"
import "math"

func lunghezza(xA, yA, xB, yB float64)(distanza float64){

    distanza = math.Sqrt( (xA - xB) * (xA - xB) + (yA - yB) * (yA - yB) )
    
    return
}

func perimetro(xA, yA, xB, yB, xC, yC float64)(ris float64){

    ris = lunghezza(xA, yA, xB, yB) + lunghezza(xB, yB, xC,yC) + lunghezza(xC, yC, xA, yA)
    
    return 
}

func main(){
// chiedo cordinate

var pxA, pxB, pxC, pyA, pyB, pyC float64
var R float64
    
    fmt.Print("inserisci dati: ")
    fmt.Scan(&pxA)
    fmt.Scan(&pyA)
    fmt.Scan(&pxB)
    fmt.Scan(&pyB)
    fmt.Scan(&pxC)
    fmt.Scan(&pyC)



    R =perimetro(pxA, pyA, pxB, pyB, pxC, pyC)
    fmt.Print("perimetro:", R)
    
}
