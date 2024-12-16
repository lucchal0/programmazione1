package main

import "fmt"

impot ("fmt")

// DICHIARAZIONI 
var s string = "Ciao, mondo!"

s := "Ciao, mondo!"

//---------------------------------------

var a [5]int
a[0] = 1
a[1] = 2

a := [5]int{1, 2, 3, 4, 5}

//---------------------------------------]]

var s []int
s = append(s, 1, 2, 3)

s := []int{1, 2, 3, 4, 5}

//---------------------------------------

var m map[string]int
m = make(map[string]int)
m["chiave"] = 42

m := map[string]int{"chiave1": 1, "chiave2": 2}

//---------------------------------------

type Persona struct {
    Nome string
    Età  int
}

p := Persona{Nome: "Mario", Età: 30}
//---------------------------------------
type Stringer interface {
    String() string
}

type Persona struct {
    Nome string
    Età  int
}

func (p Persona) String() string {
    return fmt.Sprintf("%s (%d anni)", p.Nome, p.Età)
}

//STRUTTURE DI CONTROLLO

if condizione {
    // codice se la condizione è vera
} else {
    // codice se la condizione è falsa
}

//--------------------------------------

switch espressione {
case valore1:
    // codice per valore1
case valore2:
    // codice per valore2
default:
    // codice se nessun caso è soddisfatto
}

//--------------------------------------

for i := 0; i < 10; i++ {
    // codice da eseguire
}

for condizione {
    // codice da eseguire
}

for {
    // codice da eseguire
}

for i := 0; i < 10; i++ {
    if i == 5 {
        break
    }
}

for i := 0; i < 10; i++ {
    if i%2 == 0 {
        continue
    }
    // codice da eseguire per i dispari
}

// -------------------------------------







