package main
import "fmt"

func main () {
	var sisi_atas, sisi_bawah, tinggi, luas float32
	
	sisi_atas = 12
	sisi_bawah = 8
	tinggi = 5

	luas = 0.5 * (sisi_atas + sisi_bawah) * tinggi;
	fmt.Print(luas)
}