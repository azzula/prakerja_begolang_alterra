package muncul_sekali
import (
	"fmt"
	"strconv"
)

func munculSekali(angka string) []int {
	cek := make(map[rune]int)
	list := []int{}

    for _, entry := range angka {
        cek[entry]++
    }
    for _, entry := range angka {
        if cek[entry] == 1 {
            value, err := strconv.Atoi(string(entry))
            if err != nil {
                fmt.Println("Gagal mengubah karakter ke integer:", err)
                continue
            }
            list = append(list, value)
        }
    }
    return list
}

func muncul_sekali() {
	fmt.Println(munculSekali("1234123"))	// [4]
	fmt.Println(munculSekali("76523752"))	// [6, 3]
	fmt.Println(munculSekali("12345"))		// [1 2 3 4 5]
	fmt.Println(munculSekali("1122334455"))	// []
	fmt.Println(munculSekali("0872504"))	// [8 7 2 5 4]
}