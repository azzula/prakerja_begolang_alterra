package main
import "fmt"

func main() {
	var data [10]int
	data[5] = 12
	fmt.Println(data)

	var dataSlice []int
	dataSlice = append(dataSlice, 10)
	fmt.Println(dataSlice)

	var dataBulan map[string]int = map[string]int{
		"Januari":1,
		"Desember":2,
	}
	dataBulan["Maret"] = 3
	dataBulan["Januari"] = 0
	fmt.Println(dataBulan)
}