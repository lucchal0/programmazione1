package main

import "fmt"

type Persona struct {
	nome string
	età  int
}

func main() {
	p := Persona{nome: "Luigi", età: 25}
	p = incrementa(p)
	fmt.Println("funzione: ->", p)
	p.età = 25
	fmt.Println()

	//---------------------------------------------

	fmt.Println("utilizzo metodi")
	p.incrementaEta()
	fmt.Println("motodi: ->", p)
	(&p).incrementaEta()
	fmt.Println("motodi con (&p): ->", p)
}

// funzione
func incrementa(p Persona) Persona {
	p.età++
	return p
}

//metodo

func (p *Persona) incrementaEta() {
	p.età++
}
