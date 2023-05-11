package gabung_array
import "fmt"

func ArrayMerge(arrayA, arrayB []string) []string {
	cek := make(map[string]bool)
	gabung := append(arrayA, arrayB...)
    list := []string{}

    for _, entry := range gabung {
        if _, value := cek[entry]; !value {
            cek[entry] = true
            list = append(list, entry)
        }
    }
    return list
}

func gabung_array() {
	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))
	// ["king", "devil jin", "akuma", "eddie", "steve", "geese"]

	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))
	// ["sergei", "jin", "steve", bryan]

	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa","law"}))
	// ["alisa, yoshimitsu", "devil jin", "law"]

	fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))
	// ["devil jin", "sergei"]

	fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))
	// ["hwoarang"]

	fmt.Println(ArrayMerge([]string{}, []string{}))
	// []
}