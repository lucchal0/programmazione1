package main
	
import "fmt"

func test(x, y int) (int, int) {
	return 2 * x, 2 * y
}

func main() {
	var x, y int = 10, 20
	fmt.Println(test(x, y))
}
