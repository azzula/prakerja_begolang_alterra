package main
import "fmt"

func Mapping(slice []string) map[string]int {
	cek := make(map[string]int)

 	for _, entry := range slice {
 		_, ada := cek[entry]

 		if ada {
 			cek[entry] += 1
 		} else {
 			cek[entry] = 1
 		}
 	}
 	return cek
}

func main() {
	fmt.Println(Mapping([]string{"asd", "qwe", "asd", "adi", "qwe", "qwe"}))
	// map[adi: 1 asd: 2 qwe:3]

	fmt.Println(Mapping([]string{"asd", "qwe", "asd"}))
	// map[asd: 2 qwe:1]

	fmt.Println(Mapping([]string{}))
	// map[]
}