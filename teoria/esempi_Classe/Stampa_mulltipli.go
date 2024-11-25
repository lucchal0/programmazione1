//prendere un intervallo di interi e una k stamapre i multipli di k al interno del intervallo dato 

package main 

import "fmt"

//************** SOTTOPROGRAMMA *********************
func multipli(){
var i int

    for i =1; i<=100; i++{
        if i%3 == 0{        
           fmt.Print(i," ")
        }
    }
}

//*****************************************************


func main(){

    var n int
    // stampare i multipli di 3 nel intervallo 1 a 100
    //var i int

    fmt.Print("inserire un numero:")
    fmt.Scan(&n)

    fmt.Println()
    
    multipli()

/*   
     for i =1; i<=100; i++{
        if i%3 == 0{        // 3 può essere sostituito in base a quello che ci serve ma la struttura rimane uguale
           fmt.Print(i," ") 
        }
    }

    fmt.Println()
      for i =1; i<=100; i++{
          if i%5 == 0{        
             fmt.Print(i," ")
          }
      }
                                                 Da togliere cosi da evitare la dubplicazione del codice utilizzando i sottoprogrammi
      fmt.Println()
      for i =1; i<=100; i++{
          if i%7 == 0{        
             fmt.Print(i," ")
          }
      }

       

*/ 

}


