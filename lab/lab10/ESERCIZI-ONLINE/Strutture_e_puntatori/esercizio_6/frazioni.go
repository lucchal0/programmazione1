package main

import "fmt"

type Frazione struct {
	numeratore, denominatore int
}

func NuovaFrazione(numeratore, denominatore int) *Frazione {

	//restituisce una nuova istanza del tipo Frazione inizializzata in base ai valori dei parametri numeratore e denominatore ;
	return &Frazione{numeratore: numeratore, denominatore: denominatore}
}

func String(f *Frazione) string {
	return fmt.Sprintf("%d/%d", f.numeratore, f.denominatore)
}

func main() {

	var numeratore, denominatore int

	fmt.Println("inserisci numeratore e denominatore della frazione:")
	fmt.Scan(&numeratore, &denominatore)
	fmt.Println(String(NuovaFrazione(numeratore, denominatore)))

}
