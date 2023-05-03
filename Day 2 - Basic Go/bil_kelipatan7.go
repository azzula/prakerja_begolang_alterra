package main
import "fmt"

func main () {
	for bil := 0; bil <= 70; bil++ {
		if bil % 7 == 0 {
			fmt.Println(bil)
		}
	}
}