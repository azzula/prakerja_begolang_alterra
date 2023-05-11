package main
import "fmt"

func main() {
	// function
	var hasil int
	var message string
	// hasil, message := LuasPersegiPanjang(10, 5)
	hasil, message = LuasPersegiPanjang(10, 5)
	fmt.Println(hasil, message) // 50 success

	// array => index, tipe data sejenis, ukuran tetap
	var data [5]int = [5]int{1, 5, 25}
	// data := [5]int{1, 5, 25}
	data[1] = 10
	fmt.Println(data) // [1 10 25 0 0]

	// slice => index, tipe data sejenis, ukuran dinamis
	// var dataSlice []int
	var dataSlice []int = []int{1, 2, 3, 4}
	dataSlice = append(dataSlice, 10, 20, 10)
	dataSlice[1] = 100
	fmt.Println(dataSlice) // [1 100 3 4 10 20 10]
	indexRemove := 2
	dataSlice = append(dataSlice[:indexRemove], dataSlice[indexRemove+1:]...)
	fmt.Println(dataSlice) // [1 100 4 10 20 10]

	// map => key value
	var dataBulan map[string]int = map[string]int{
		"januari":1,
		"desember":12,
	}
	dataBulan["maret"] = 3
	dataBulan["januari"] = 0
	// result := dataBulan["januari"]
	// _, isExist := dataBulan["januari"]
	fmt.Println(dataBulan) // map[desember:12 januari:0 maret:3]
}

// function
func LuasPersegiPanjang(p int, l int) (int, string) {
	L := p * l
	return L, "success"
}