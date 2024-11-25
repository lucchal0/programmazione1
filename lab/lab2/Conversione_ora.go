package main

import "fmt"

func main() {
	var scelta float64
	var secondi float64
	var minuti float64
	var ore float64
	var giorni float64
	var anni float64

	fmt.Println("scelta: 1) secondi -> ore 2) secondi -> minuti 3) minuti -> ore 4) minuti -> secondi 5) ore -> secondi 6) ore -> minuti 7) minuti -> giorni e ore 8) minuti -> anni e giorni")
	fmt.Scan(&scelta)

	switch scelta {
	case 1:
		//secondi -> ore
		fmt.Print("inserire i secondi da convertire in ore: ")
		fmt.Scan(&secondi)

		fmt.Println(secondi, "secondi corrispondono a", secondi/3600, "ore")
		break

	case 2:
		//secondi -> minuti
		fmt.Print("inserire i secondi da convertire in minuti: ")
		fmt.Scan(&secondi)

		fmt.Println(secondi, "secondi corrispondono a", secondi/60, "minuti")
		break

	case 3:
		//minuti -> ore
		fmt.Print("inserire i minuti da convertire in ore: ")
		fmt.Scan(&minuti)

		fmt.Println(minuti, "minuti corrispondono a", minuti/60, "ore")
		break

	case 4:
		//minuti -> secondi
		fmt.Print("inserire i minuti da convertire in secondi: ")
		fmt.Scan(&minuti)

		fmt.Println(minuti, "minuti corrispondono a", minuti*60, "secondi")
		break

	case 5:
		//ore -> secondi
		fmt.Print("inserire le ore da convertire in secondi: ")
		fmt.Scan(&ore)

		fmt.Println(ore, "ore corrispondono a", ore*3600, "secondi")
		break

	case 6:
		//ore -> minuti
		fmt.Print("inserire le ore da convertire in minuti: ")
		fmt.Scan(&ore)

		fmt.Println(ore, "ore corrispondono a", ore*60, "minuti")
		break

	case 7:
		//minuti -> giorni e ore
		fmt.Print("inserire i minuti da convertire in giorni e ore: ")
		fmt.Scan(&minuti)

		// Conversione in int per usare l'operatore modulo
		giorni = float64(int(minuti) / 1440)
		ore = float64((int(minuti) % 1440) / 60)
		fmt.Println(minuti, "minuti corrispondono a", giorni, "giorni e", ore, "ore")
		break

	case 8:
		// minuti -> anni e giorni
		fmt.Print("inserire i minuti da convertire in anni e giorni: ")
		fmt.Scan(&minuti)

		// Conversione in int per usare l'operatore modulo
		anni = float64(int(minuti) / 525600)
		giorni = float64((int(minuti) % 525600) / 1440)
		fmt.Println(minuti, "minuti corrispondono a", anni, "anni e", giorni, "giorni")
		break
	}
}
