package main

import "fmt"

// stampare una scacchiera N X N

// B -> bianco N -> nero

func main(){

var riga int
var colonne int
var i,j int

      fmt.Print("inserisci quanto grande le righe:")
      fmt.Scan(&riga)

      fmt.Print("colonna:")
      fmt.Scan(&colonne)

      for i= 0; i < N; i++ {
          for j= 0; j< colonne; j++{

                if i % 2 == j % 2{
                    fmt.Print('B')
                }else{
                    fmt.Print('N')
                }
          } 
          fmt.Print('\n')
      } 

}
