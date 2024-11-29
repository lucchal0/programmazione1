// stampare una scacchiera NxM dove M e N sono forniti dall'utente da riga di comando

package main

import (
	"fmt"
	"os"
	"strconv"
)

func inputDimensione(argomenti []string)(N,M int){
	N= strconv.Atoi(argomenti[1])
	M= strconv.Atoi(argomenti[2])
}

//input: numero di righe N e il numero di colonne M
// output : S [][]bool, scacchiera di NxM, false = nero, true = bianco
func creaScacchiera(M,N int)(S[][]bool){
	S = make([][]bool, N)
	for i:=0; i<N; i++{
		S[i]= make([]bool, M)
	}
	for i:=0; i<N; i++{
		for j:=0; j<M; i++{
			if(i%2) == (j%2){
				S[i][j] = true
			}else{
				S[i][j] = false
			}
		}
	}
	return 
}

func stampascacchiera(){

}

func main(){
var N, M int
var T [][] bool
	//input 
	N,M=inputDimensione(os.Args)


	//creare scacchiera
	T = creaScacchiera(N,M)

	//stampare
	stampa
}