package main
import "fmt"

func main() {
	// var name string = "Sufyan"
	// var age int = 26
	// var age2 uint = 7
	// var phi float32 = 3.14
	// var isMan bool = true

	for i := 0; i <= 3; i++ {
		for j := 0; j <= 3; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}