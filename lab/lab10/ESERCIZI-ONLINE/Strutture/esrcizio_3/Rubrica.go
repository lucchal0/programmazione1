package main

import (
	"fmt"
)

type indirizzo struct {
	via          string
	numeroCivico uint
	cap          string
	citta        string
}

type contatto struct {
	cognome   string
	nome      string
	telefono  string
	indirizzo indirizzo
}

type rubrica []contatto

func main() {

	adress := indirizzo{"Celoria", 18, "26133", "Milano"}
	// fmt.Printf("%#v\n", adress)

	contact := contatto{"Rossi", "Mario", "03030030302", adress}
	// fmt.Printf("%#v\n", contact)

	r := rubrica{contact}

	fmt.Printf("%#v", r)
}

func NuovaRubrica() rubrica {
	return rubrica{}
}

func InserisciContatto(r rubrica, cognome, nome string, via string, numero uint, cap, citta string, telefono string) rubrica {

	for _, c:=range r{
		if c.cognome == cognome && c.nome == nome{
			return r
		}
	}

	adress := indirizzo{via, numero, cap, citta}
	contact := contatto{cognome, nome, telefono, adress}

	r = append(r, contact)
	return r
}

func eliminaContatto(r rubrica, cognome,nome string) rubrica{
	// NuovaRubrica := NuovaRubrica()
	// for _, c:=range r{
	// 	if c.cognome != cognome || c.nome != nome{
	// 		NuovaRubrica = append(NuovaRubrica, c)
	// 	}
	// }
	// return r

	for i, c:= range r{
		if c.cognome == cognome || c.nome == nome{
			r = append(r[:i], r[i+1:]... )
			break
		}
	}
	return r
}

func stampaContatto(c contatto){
	fmt.Printf("%s %s: %s %d, %s, %s - Tel. %s\n", c.nome,c.cognome,c.indirizzo.via,c.indirizzo.numeroCivico, c.indirizzo.citta, c.indirizzo.cap, c.telefono)
}

func stampaRubrica(r rubrica){
	fmt.Println("Rubrica:")
	for i:=range r{
		fmt.Printf("[%d] - ", i+1)
		stampaContatto(r[i])
	}
}

func aggiornaContatto()
