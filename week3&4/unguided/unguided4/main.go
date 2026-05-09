package main

import "fmt"

func main() {
	var pilihan int
	fmt.Println("---PROGRAM BANGUN DATAR---")
	fmt.Println("1.hitung luas persegi ")
	fmt.Println("2.hitung luas persegi panjang")
	fmt.Println("3.hitung luas lingkaran")
	fmt.Print("masukan pilihan:")
	fmt.Scan(&pilihan)
	switch pilihan {
	case 1:
		persegi()
	case 2:
		persegiPanjang()
	case 3:
		lingkaran()

	}
}

func persegi() {
	var sisi int
	fmt.Print("masukan sisi: ")
	fmt.Scan(&sisi)
	luas := sisi * sisi
	keliling := 4 * sisi
	fmt.Println("luas persegi:", luas)
	fmt.Println("luas keliling:", keliling)
}

func persegiPanjang() {
	var panjang, lebar int
	fmt.Print("masukan panjang:")
	fmt.Scan(&panjang)
	fmt.Print("masukan lebar:")
	fmt.Scan(&lebar)
	luas := panjang * lebar
	keliling := 2 * (panjang + lebar)
	fmt.Println("luas persegi panjang:", luas)
	fmt.Println("luas persegi keliling:", keliling)
}

func lingkaran() {
	var jari float32
	fmt.Print("masukan jari-jari:")
	fmt.Scan(&jari)
	luas := 3.14 * jari * jari
	keliling := 2 * 3.14 * jari
	fmt.Println("luas lingkaran:", luas)
	fmt.Println("keliling lingkaran:", keliling)

}
