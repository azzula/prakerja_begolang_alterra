package main
import (
	"fmt"
	"prakerja_begolang_alterra/Day4_Advance_Go/featurea"
	"prakerja_begolang_alterra/Day4_Advance_Go/featureb"
)

func main() {
	// pointer => menangkap alamat memori
	var a int
	var aPointer *int = &a
	fmt.Println(&a) // 0x...
	fmt.Println(aPointer) // 0x...

	var b int = 10
	gantiNilai(&b) // ditimpa 50
	b = gantiNilai2(b) // ditimpa 100
	fmt.Println(b) // 100 (nilai terakhir)

	// struct => object
	var m Mobil = Mobil{mesin: 10.1, stir: 20}
	// var m Mobil = Mobil{10, 20, "c", 40}
	// var m Mobil
	// m.ban = 10
	// m.mesin = 20
	panggil(m)

	// method
	var mobil Mobil = Mobil{ban: 5}
	mobil.jalan()
	fmt.Println(mobil) // {15 0 "" 0}

	// interface
	var myRestoran = Restoran{20, 10}
	var myMenu MenuRestoran
	myMenu = myRestoran
	fmt.Println("Nasi Goreng ", myMenu.buatNasgor()) // 30
	fmt.Println("Mie Goreng ", myMenu.buatMiegor()) // 20

	// dinamis
	tampil(10)
	tampil("10")

	// error catch
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	panic("Error A")

	// package
	featurea.TampilRestoran() // <-- unreachable code
	featureb.TampilMobil() // <-- unreachable code
}

// pointer
func gantiNilai(b *int) {
	fmt.Println(*b) // 10 (nilai asli)
	*b = 50
}

func gantiNilai2(a int) int {
	return 100
}

// struct
type Mobil struct {
	ban		int
	stir	int
	knalpot	string
	mesin	float32
}

func panggil(mobil Mobil) {
	// fmt.Println(mobil.ban)
	fmt.Println(mobil) // {0 20 "" 10.1}
}

// method
func (mobil *Mobil) jalan() int {
	fmt.Println("Ban method ", mobil) // {5 0 "" 0}
	mobil.ban = mobil.ban + 10
	return mobil.ban
}

// interface
type Restoran struct {
	nasgor int
	miegor int
}

func (r Restoran) buatNasgor() int {
	r.nasgor = r.nasgor + 10
	return r.nasgor
}

func (r Restoran) buatMiegor() int {
	r.miegor = r.miegor + 10
	return r.miegor
}

type MenuRestoran interface {
	buatNasgor() int
	buatMiegor() int
}

// dinamis
func tampil(data interface{}) {
	fmt.Println(data) // 10 // "10"
}
