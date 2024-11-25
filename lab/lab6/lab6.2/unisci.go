package main

import "bufio"
import "fmt"
import "os"

func main() {

	var newStr string
	for {
		fmt.Println("inserisci una stringa:")
		s := bufio.NewScanner(os.Stdin)
		s.Scan()
		if s.Text() == "" {
			fmt.Print("stringa completa:", newStr)
			break
		}
		newStr = newStr + s.Text() + " "

	}
}
